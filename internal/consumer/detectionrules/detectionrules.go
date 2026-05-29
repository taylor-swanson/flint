// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package detectionrules

import (
	"fmt"
	"go/token"
	"log/slog"
	"os"
	"strings"

	"github.com/taylor-swanson/flint/internal/consumer"
	"github.com/taylor-swanson/flint/internal/field"
	"github.com/taylor-swanson/flint/internal/file"
)

const Name = "detectionrules"

var _ consumer.Consumer = new(detectionrules)

type detectionrules struct {
	files []string
}

func (c *detectionrules) Consume(state *consumer.State) ([]field.Usage, error) {
	var usages []field.Usage

	for _, filename := range c.files {
		f, err := file.Load(filename)
		if err != nil {
			return nil, err
		}
		indexPatterns, fields, err := parse(f.Bytes())
		if err != nil {
			return nil, err
		}

		for _, indexPattern := range indexPatterns {
			for _, v := range fields {
				usage, err := state.FieldMap.Lookup(v, indexPattern)
				if err != nil {
					slog.Warn("Unable to lookup definition", slog.String("field", v), slog.String("index_pattern", indexPattern))
					continue
				}
				// Note: exact position information not yet available.
				usage.Position = token.Position{Filename: filename}

				usages = append(usages, usage)
			}
		}
	}

	return usages, nil
}

func New(_ string, files []string) (consumer.Consumer, error) {
	var c detectionrules

	for _, f := range files {
		if stat, err := os.Stat(f); err == nil && !stat.IsDir() {
			if strings.HasSuffix(f, ".toml") {
				c.files = append(c.files, f)
			}
		}
	}

	if len(c.files) == 0 {
		return nil, fmt.Errorf("%s: no files found", Name)
	}

	return &c, nil
}

func init() {
	consumer.Register(Name, New)
}
