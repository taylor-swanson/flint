// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package integrations

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
	"go.yaml.in/yaml/v4"

	"github.com/taylor-swanson/flint/internal/consumer"
	"github.com/taylor-swanson/flint/internal/consumer/generic"
	"github.com/taylor-swanson/flint/internal/field"
)

const Name = "integrations"

type manifest struct {
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`
	Version string `yaml:"version"`
}

type dataStreamManifest struct {
	Dataset string `yaml:"dataset"`
	Type    string `yaml:"type"`
}

var dataStreamPatterns = []string{
	"elasticsearch/ingest_pipeline/*.yml",
}

var generalPatterns = []string{
	"elasticsearch/ingest_pipeline/*.yml",
	"kibana/dashboard/*.json",
	"kibana/search/*.json",
}

var _ consumer.Consumer = new(integrations)

type integrations struct {
	pkgDirs []string
}

func (c *integrations) Consume(state *consumer.State) ([]field.Usage, error) {
	var usages []field.Usage
	var err error

	for _, pkgDir := range c.pkgDirs {
		var m manifest
		if err = parseYAMLFile(filepath.Join(pkgDir, "manifest.yml"), &m); err != nil {
			return nil, fmt.Errorf("failed to load package manifest: %w", err)
		}

		generalIndexPattern := "logs-" + m.Name + "*"
		for _, pattern := range generalPatterns {
			matches, err := doublestar.FilepathGlob(filepath.Join(pkgDir, pattern))
			if err != nil {
				return nil, err
			}

			for _, filename := range matches {
				got, err := generic.Parse(state.FieldMap, generalIndexPattern, filename)
				if err != nil {
					return nil, err
				}

				usages = append(usages, got...)
			}
		}

		var dataStreamDirs []string
		if dataStreamDirs, err = filepath.Glob(filepath.Join(pkgDir, "data_stream", "*")); err != nil {
			return nil, fmt.Errorf("failed to glob data_stream directories: %w", err)
		}
		for _, dataStreamDir := range dataStreamDirs {
			if stat, err := os.Stat(dataStreamDir); err != nil || !stat.IsDir() {
				continue
			}

			var dsm dataStreamManifest
			if err = parseYAMLFile(filepath.Join(dataStreamDir, "manifest.yml"), &dsm); err != nil {
				return nil, fmt.Errorf("failed to load data_stream manifest: %w", err)
			}
			dataset := dsm.Dataset
			if dsm.Dataset == "" {
				dataset = m.Name + "." + filepath.Base(dataStreamDir)
			}
			datasetType := dsm.Type
			if datasetType == "" {
				datasetType = "logs"
			}

			dsIndexPattern := datasetType + "-" + dataset + "-*"
			for _, pattern := range dataStreamPatterns {
				matches, err := doublestar.FilepathGlob(filepath.Join(dataStreamDir, pattern))
				if err != nil {
					return nil, err
				}

				for _, filename := range matches {
					got, err := generic.Parse(state.FieldMap, dsIndexPattern, filename)
					if err != nil {
						return nil, err
					}

					usages = append(usages, got...)
				}
			}
		}
	}

	return usages, nil
}

func New(_ string, paths []string) (consumer.Consumer, error) {
	var c integrations

	for _, f := range paths {
		if stat, err := os.Stat(f); err == nil && stat.IsDir() {
			if _, err = os.Stat(filepath.Join(f, "manifest.yml")); err == nil {
				c.pkgDirs = append(c.pkgDirs, f)
			}
		}
	}

	if len(c.pkgDirs) == 0 {
		return nil, fmt.Errorf("%s: no package directories found", Name)
	}

	return &c, nil
}

func init() {
	consumer.Register(Name, New)
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
