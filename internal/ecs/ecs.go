// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package ecs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"go.yaml.in/yaml/v4"

	"github.com/taylor-swanson/flint/internal/field"
)

func NewResolver(ref, cacheDir string) (*Resolver, error) {
	if r, err := loadResolver(ref, cacheDir); err == nil {
		return r, nil
	}

	r, err := buildResolver(ref)
	if err != nil {
		return nil, err
	}

	if err = saveResolver(ref, cacheDir, r); err != nil {
		slog.Error("Failed to save resolver", slog.String("error", err.Error()))
	}

	return r, nil
}

func LookupVersion(ref, cacheDir string) (string, error) {
	r, err := NewResolver(ref, cacheDir)
	if err != nil {
		return "", err
	}

	return r.version, nil
}

type Resolver struct {
	version string
	ref     string
	fields  map[string]*field.Field
}

func (r *Resolver) Lookup(name string) *field.Field {
	return r.fields[name]
}

func (r *Resolver) Version() string {
	return r.version
}

func (r *Resolver) Ref() string {
	return r.ref
}

type resolverFile struct {
	Version string                  `json:"version"`
	Ref     string                  `json:"ref"`
	Fields  map[string]*field.Field `json:"fields"`
}

func (r *Resolver) UnmarshalJSON(bytes []byte) error {
	var rf resolverFile
	if err := json.Unmarshal(bytes, &rf); err != nil {
		return err
	}

	r.version = rf.Version
	r.ref = rf.Ref
	r.fields = rf.Fields

	return nil
}

func (r *Resolver) MarshalJSON() ([]byte, error) {
	return json.Marshal(&resolverFile{
		Version: r.version,
		Ref:     r.ref,
		Fields:  r.fields,
	})
}

func getCacheFilename(ref, cacheDir string) string {
	return filepath.Join(cacheDir, "ecs-"+ref+".json")
}

func loadResolver(ref, cacheDir string) (*Resolver, error) {
	cacheFilename := getCacheFilename(ref, cacheDir)

	raw, err := os.ReadFile(cacheFilename)
	if err != nil {
		return nil, err
	}

	var r Resolver
	if err = json.Unmarshal(raw, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func saveResolver(ref, cacheDir string, r *Resolver) error {
	cacheFilename := getCacheFilename(ref, cacheDir)

	if err := os.MkdirAll(filepath.Dir(cacheFilename), 0o755); err != nil {
		return err
	}

	data, err := json.Marshal(r)
	if err != nil {
		return err
	}
	if err = os.WriteFile(cacheFilename, data, 0o644); err != nil {
		return err
	}

	return nil
}

type fieldDef struct {
	Type        field.Type `yaml:"type"`
	Description string     `yaml:"description"`
	Normalize   any        `yaml:"normalize"`
}

func buildResolver(ref string) (*Resolver, error) {
	isArray := func(v any) bool {
		switch t := v.(type) {
		case string:
			if t == "array" {
				return true
			}
		case []any:
			for _, tv := range t {
				switch vv := tv.(type) {
				case string:
					if vv == "array" {
						return true
					}
				}
			}
		}

		return false
	}

	ref = strings.TrimPrefix(ref, "git@")

	versionRaw, err := fetchFile("https://raw.githubusercontent.com/elastic/ecs/" + ref + "/version")
	if err != nil {
		return nil, err
	}

	raw, err := fetchFile("https://raw.githubusercontent.com/elastic/ecs/" + ref + "/generated/ecs/ecs_flat.yml")
	if err != nil {
		return nil, err
	}
	var fieldDefs map[string]fieldDef
	if err = yaml.Unmarshal(raw, &fieldDefs); err != nil {
		return nil, err
	}

	r := Resolver{
		ref:     ref,
		version: string(bytes.TrimSpace(versionRaw)),
		fields:  make(map[string]*field.Field),
	}

	for k, v := range fieldDefs {
		r.fields[k] = &field.Field{
			Name:  k,
			Type:  v.Type,
			Kind:  field.KindECS,
			Array: isArray(v.Normalize),
		}

		// Add lat, long, type, and coordinates sub-fields for geo_point types.
		if v.Type == field.TypeGeoPoint {
			name := k + ".lat"
			r.fields[name] = &field.Field{
				Name: name,
				Type: field.TypeGeoPoint,
				Kind: field.KindECS,
			}
			name = k + ".lon"
			r.fields[name] = &field.Field{
				Name: name,
				Type: field.TypeGeoPoint,
				Kind: field.KindECS,
			}
			name = k + ".type"
			r.fields[name] = &field.Field{
				Name: name,
				Type: field.TypeGeoPoint,
				Kind: field.KindECS,
			}
			name = k + ".coordinates"
			r.fields[name] = &field.Field{
				Name:  name,
				Type:  field.TypeGeoPoint,
				Kind:  field.KindECS,
				Array: true,
			}
		}
	}

	return &r, nil
}

func fetchFile(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get remote file %q: %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get remote file %q, status code %d", url, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body from %q: %w", url, err)
	}

	return body, nil
}
