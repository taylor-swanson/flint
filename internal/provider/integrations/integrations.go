// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package integrations

import (
	"errors"
	"fmt"
	"go/token"
	"log/slog"
	"os"
	"path/filepath"

	"go.yaml.in/yaml/v4"

	"github.com/taylor-swanson/flint/internal/ecs"
	"github.com/taylor-swanson/flint/internal/field"
	"github.com/taylor-swanson/flint/internal/file"
	"github.com/taylor-swanson/flint/internal/provider"
)

const (
	Name = "integrations"
)

type manifest struct {
	Name       string `yaml:"name"`
	Type       string `yaml:"type"`
	Version    string `yaml:"version"`
	Conditions struct {
		Kibana struct {
			Version string `yaml:"version"`
		} `yaml:"kibana"`
	} `yaml:"conditions"`
}

type buildManifest struct {
	Dependencies struct {
		ECS struct {
			Reference string `yaml:"reference"`
		} `yaml:"ecs"`
	} `yaml:"dependencies"`
}

type dataStreamManifest struct {
	Dataset string `yaml:"dataset"`
	Type    string `yaml:"type"`
}

type pkgInfo struct {
	Name    string
	Title   string
	Type    string
	ECSRef  string
	Version string
	Stack   string
}

type fieldDef struct {
	Name        string     `yaml:"name"`
	External    string     `yaml:"external"`
	Path        *string    `yaml:"path"`
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

func parseYAMLFile(filename string, v any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read YAML file %q: %w", filename, err)
	}
	if err = yaml.Unmarshal(data, v); err != nil {
		return fmt.Errorf("failed to unmarshal YAML file %q: %w", filename, err)
	}

	return nil
}

func loadFields(dirpath string, ecsResolver *ecs.Resolver, fieldMap *field.Map, index *field.Index) error {
	var definitions []field.Definition

	files, err := filepath.Glob(filepath.Join(dirpath, "fields", "*.yml"))
	if err != nil {
		return err
	}

	var flatten func(def fieldDef, parent string)
	flatten = func(def fieldDef, parent string) {
		name := field.Join(parent, def.Name)

		if def.Type != field.TypeGroup {
			if ecsResolver != nil && def.External == "ecs" {
				ecsField := ecsResolver.Lookup(name)
				if ecsField != nil {
					definitions = append(definitions, field.Definition{
						Field:    ecsField,
						Index:    index,
						Position: def.Position,
					})
				}
			} else {
				definitions = append(definitions, field.Definition{
					Field: &field.Field{
						Name:   name,
						Target: def.Path,
						Kind:   field.KindVendor,
						Type:   def.Type,
					},
					Index:    index,
					Position: def.Position,
				})
			}
		}

		for _, v := range def.Fields {
			flatten(v, name)
		}
		for _, v := range def.MultiFields {
			flatten(v, name)
		}
	}

	for _, f := range files {
		var fieldDefs []fieldDef
		if err = parseYAMLFile(f, &fieldDefs); err != nil {
			return err
		}

		file.Annotate(f, &fieldDefs)

		for _, def := range fieldDefs {
			flatten(def, "")
		}
	}

	fieldMap.AddDefinition(definitions...)

	return nil
}

func getPackageFields(pkgDir string, state *provider.State) error {
	var err error

	var m manifest
	if err = parseYAMLFile(filepath.Join(pkgDir, "manifest.yml"), &m); err != nil {
		return fmt.Errorf("failed to load package manifest: %w", err)
	}

	var bm buildManifest
	if err = parseYAMLFile(filepath.Join(pkgDir, "_dev", "build", "build.yml"), &bm); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to load build manifest: %w", err)
	}

	pkg := pkgInfo{
		Name:    m.Name,
		Type:    m.Type,
		ECSRef:  bm.Dependencies.ECS.Reference,
		Version: m.Version,
		Stack:   m.Conditions.Kibana.Version,
	}

	var ecsResolver *ecs.Resolver
	if pkg.ECSRef != "" {
		if ecsResolver, err = ecs.NewResolver(pkg.ECSRef, state.CacheDir); err != nil {
			slog.Warn("Failed to create ECS resolver", slog.String("package", pkg.Name), slog.String("ecs_ref", pkg.ECSRef), slog.String("error", err.Error()))
		}
	}

	if pkg.Type == "input" {
		dataset := pkg.Name + ".generic"

		index := &field.Index{
			Name:    dataset,
			Pattern: "logs-" + dataset + "-*",
			ECSRef:  pkg.ECSRef,
			Version: pkg.Version,
			Stack:   pkg.Stack,
		}
		if err = loadFields(pkgDir, ecsResolver, state.FieldMap, index); err != nil {
			return fmt.Errorf("failed to load input fields: %w", err)
		}
	} else {
		var dataStreamDirs []string
		if dataStreamDirs, err = filepath.Glob(filepath.Join(pkgDir, "data_stream", "*")); err != nil {
			return fmt.Errorf("failed to glob data_stream directories: %w", err)
		}

		for _, dataStreamDir := range dataStreamDirs {
			if stat, err := os.Stat(dataStreamDir); err != nil || !stat.IsDir() {
				continue
			}

			var dsm dataStreamManifest
			if err = parseYAMLFile(filepath.Join(dataStreamDir, "manifest.yml"), &dsm); err != nil {
				return fmt.Errorf("failed to load data_stream manifest: %w", err)
			}
			dataset := dsm.Dataset
			if dsm.Dataset == "" {
				dataset = pkg.Name + "." + filepath.Base(dataStreamDir)
			}
			datasetType := dsm.Type
			if datasetType == "" {
				datasetType = "logs"
			}

			index := &field.Index{
				Name:    dataset,
				Pattern: datasetType + "-" + dataset + "-*",
				ECSRef:  pkg.ECSRef,
				Version: pkg.Version,
				Stack:   pkg.Stack,
			}
			if err = loadFields(dataStreamDir, ecsResolver, state.FieldMap, index); err != nil {
				return fmt.Errorf("failed to load data stream fields: %w", err)
			}
		}
	}

	return nil
}

var _ provider.Provider = new(integrations)

type integrations struct {
	pkgDirs []string
}

func (p *integrations) Provide(state *provider.State) error {
	for _, pkg := range p.pkgDirs {
		if err := getPackageFields(pkg, state); err != nil {
			return err
		}
	}

	return nil
}

func New(path string) (provider.Provider, error) {
	var p integrations

	if _, err := os.Stat(filepath.Join(path, "manifest.yml")); err == nil {
		p.pkgDirs = append(p.pkgDirs, path)
	} else {
		{ // Check for a packages directory (i.e., if this is the integrations repo).
			packagesDir := filepath.Join(path, "packages")
			if stat, err := os.Stat(packagesDir); err == nil && stat.IsDir() {
				path = packagesDir
			}
		}

		matches, err := filepath.Glob(filepath.Join(path, "*"))
		if err != nil {
			return nil, fmt.Errorf("%s: failed to glob: %w", Name, err)
		}
		for _, v := range matches {
			if _, err = os.Stat(filepath.Join(v, "manifest.yml")); err == nil {
				p.pkgDirs = append(p.pkgDirs, v)
			}
		}
	}

	if len(p.pkgDirs) == 0 {
		return nil, fmt.Errorf("%s: no packages found in %q", Name, path)
	}

	return &p, nil
}

func init() {
	provider.Register(Name, New)
}
