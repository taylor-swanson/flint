// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package generic

import (
	"log/slog"
	"strings"

	"github.com/taylor-swanson/flint/internal/field"
	"github.com/taylor-swanson/flint/internal/file"
)

type match struct {
	Start int
	End   int
}

var replacer = strings.NewReplacer("ctx.", "", "?", "")

func Parse(fieldMap *field.Map, indexPattern, filename string) ([]field.Usage, error) {
	f, err := file.Load(filename)
	if err != nil {
		return nil, err
	}

	var usages []field.Usage
	matches := parseFields(f.Bytes())
	for _, m := range matches {
		value := string(f.Bytes()[m.Start:m.End])
		value = replacer.Replace(value)

		usage, err := fieldMap.Lookup(value, indexPattern)
		if err != nil {
			slog.Error("Failed to lookup field", slog.String("field", value), slog.String("index_pattern", indexPattern), slog.String("error", err.Error()))
			continue
		}
		usage.Position = f.Position(m.Start)

		usages = append(usages, usage)
	}

	return usages, nil
}
