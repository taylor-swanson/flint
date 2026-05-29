// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package beats

import (
	"fmt"
	"go/token"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"go.yaml.in/yaml/v4"

	"github.com/taylor-swanson/flint/internal/ecs"
	"github.com/taylor-swanson/flint/internal/field"
	"github.com/taylor-swanson/flint/internal/file"
	"github.com/taylor-swanson/flint/internal/provider"
)

const (
	Name = "beats"
)

var _ provider.Provider = new(beats)

const DefaultECSRef = "v8.0.0" // Note: Figure out how to read this from the repo.

type fieldDef struct {
	Name        string     `yaml:"name"`
	Type        field.Type `yaml:"type"`
	Fields      []fieldDef `yaml:"fields"`
	MultiFields []fieldDef `yaml:"multi_fields"`

	Position    token.Position `json:"-" yaml:"-"`
	JSONPointer string         `json:"-" yaml:"-"`
}

func (v *fieldDef) UnmarshalYAML(node *yaml.Node) error {
	type f fieldDef
	x := (*f)(v)
	if err := node.Decode(x); err != nil {
		return err
	}

	v.Position.Line = node.Line
	v.Position.Column = node.Column

	return nil
}

var beatNames = []string{
	"auditbeat",
	"filebeat",
	"heartbeat",
	"metricbeat",
	"osquerybeat",
	"packetbeat",
	"winlogbeat",
}

var versionRE = regexp.MustCompile(`const defaultBeatVersion = "(\d+.\d+.\d+(?:-[a-zA-Z0-9]+)?)"`)

func readVersion(root string) (string, error) {
	raw, err := os.ReadFile(filepath.Join(root, "libbeat", "version", "version.go"))
	if err != nil {
		return "", fmt.Errorf("failed to read libbeat version: %w", err)
	}

	got := versionRE.FindSubmatch(raw)

	if len(got) == 0 {
		return "", fmt.Errorf("failed to find libbeat version")
	}

	return string(got[1]), nil
}

type beats struct {
	path string
}

func (p *beats) parseFields(filename string, ecsResolver *ecs.Resolver) ([]field.Definition, error) {
	var definitions []field.Definition

	var flatten func(def fieldDef, parent string)
	flatten = func(def fieldDef, parent string) {
		name := def.Name
		if parent != "" {
			name = parent + "." + name
		}

		isArray := false
		if name != "" && def.Type != field.TypeGroup {
			kind := field.KindVendor
			if def.Type != field.TypeAlias {
				if ecsField := ecsResolver.Lookup(name); ecsField != nil {
					if ecsField.Type == def.Type {
						kind = field.KindECS
					} else {
						slog.Warn("ECS field type mismatch", slog.String("field", name), slog.String("expected", string(ecsField.Type)), slog.String("observed", string(def.Type)), slog.String("filename", filename))
					}
					isArray = ecsField.Array
				}
			}

			definitions = append(definitions, field.Definition{
				Field: &field.Field{
					Name:  name,
					Type:  def.Type,
					Kind:  kind,
					Array: isArray,
				},
				Position: def.Position,
			})
		}

		for _, v := range def.MultiFields {
			flatten(v, name)
		}
		for _, v := range def.Fields {
			flatten(v, name)
		}
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var fieldDefs []fieldDef
	if err = yaml.Unmarshal(data, &fieldDefs); err != nil {
		return nil, err
	}

	file.Annotate(filename, &fieldDefs)

	for _, f := range fieldDefs {
		flatten(f, "")
	}

	return definitions, nil
}

func (p *beats) getFiles(beatName string) ([]string, error) {
	var files []string

	dirs := []string{
		filepath.Join(p.path, beatName),
		filepath.Join(p.path, "x-pack", beatName),
	}

	for _, dir := range dirs {
		matches, err := doublestar.FilepathGlob(filepath.Join(dir, "**/_meta/fields*.yml"))
		if err != nil {
			return nil, fmt.Errorf("failed to list field files: %w", err)
		}
		matches = slices.DeleteFunc(matches, func(s string) bool {
			return strings.Contains(s, "testdata") || strings.Contains(s, "scripts")
		})

		files = append(files, matches...)
	}

	return files, nil
}

func (p *beats) gatherLibbeatFields(ecsResolver *ecs.Resolver) ([]field.Definition, error) {
	files, err := p.getFiles("libbeat")
	if err != nil {
		return nil, err
	}

	var fields []field.Definition
	for _, f := range files {
		got, err := p.parseFields(f, ecsResolver)
		if err != nil {
			return nil, fmt.Errorf("failed to parse fields file %q: %w", f, err)
		}
		fields = append(fields, got...)
	}

	slog.Debug("Found libbeat fields", slog.Int("fields", len(fields)), slog.Int("files", len(files)))

	return fields, nil
}

func (p *beats) gatherBeatFields(fieldMap *field.Map, ecsResolver *ecs.Resolver, libbeatFields []field.Definition, beat, version string) error {
	files, err := p.getFiles(beat)
	if err != nil {
		return err
	}

	definitions := slices.Clone(libbeatFields)
	for _, f := range files {
		got, err := p.parseFields(f, ecsResolver)
		if err != nil {
			return fmt.Errorf("failed to parse fields file %q: %w", f, err)
		}
		definitions = append(definitions, got...)
	}

	index := &field.Index{
		Name:    beat,
		Pattern: beat + "-*",
		Version: version,
		Stack:   "^" + version,
	}
	for i := range definitions {
		definitions[i].Index = index
	}

	fieldMap.AddDefinition(definitions...)

	return nil
}

func (p *beats) Provide(state *provider.State) error {
	ecsResolver, err := ecs.NewResolver(DefaultECSRef, state.CacheDir)
	if err != nil {
		return fmt.Errorf("%s: failed to create ecs resolver: %w", Name, err)
	}

	version, err := readVersion(p.path)
	if err != nil {
		return fmt.Errorf("%s: failed to read version: %w", Name, err)
	}

	libbeatFields, err := p.gatherLibbeatFields(ecsResolver)
	if err != nil {
		return err
	}

	for _, beatName := range beatNames {
		if err = p.gatherBeatFields(state.FieldMap, ecsResolver, libbeatFields, beatName, version); err != nil {
			return fmt.Errorf("failed to get beat fields: %w", err)
		}
	}

	return nil
}

func New(path string) (provider.Provider, error) {
	p := beats{path: path}

	if _, err := os.Stat(p.path); err != nil {
		return nil, fmt.Errorf("%s: failed to locate directory %q: %w", Name, p.path, err)
	}

	return &p, nil
}

func init() {
	provider.Register(Name, New)
}
