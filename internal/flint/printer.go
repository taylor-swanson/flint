// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package flint

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/fatih/color"
)

func PrintJSON(w io.Writer, report *Report, minify bool) error {
	enc := json.NewEncoder(w)
	if !minify {
		enc.SetIndent("", "  ")
	}
	enc.SetEscapeHTML(false)

	return enc.Encode(report)
}

func PrintText(w io.Writer, report *Report, noColor bool) error {
	red := color.New(color.FgRed)
	yellow := color.New(color.FgYellow)
	green := color.New(color.FgGreen)
	bold := color.New(color.Bold)

	if noColor {
		red.DisableColor()
		yellow.DisableColor()
		green.DisableColor()
		bold.DisableColor()
	}

	var err error
	if _, err = bold.Fprint(w, "Time: "); err != nil {
		return err
	}
	if _, err = fmt.Fprintln(w, report.Timestamp.Format(time.RFC3339)); err != nil {
		return err
	}

	if _, err = bold.Fprintln(w, "Files"); err != nil {
		return err
	}
	for _, filename := range report.Files {
		if _, err = bold.Fprintf(w, "-- %s %s\n", filename, strings.Repeat("-", max(76-len(filename), 0))); err != nil {
			return err
		}

		fr := report.FileReports[filename]
		for _, indexName := range fr.Indices {
			ir := fr.IndexReports[indexName]
			if _, err = bold.Fprintln(w, "  -", indexName); err != nil {
				return err
			}
			for _, usage := range ir.Usages {
				if _, err = bold.Fprintln(w, "    -", usage.Name); err != nil {
					return err
				}
				if usage.Position.IsValid() {
					if _, err = fmt.Fprintf(w, "      %s:%d:%d\n", usage.Position.Filename, usage.Position.Line, usage.Position.Column); err != nil {
						return err
					}
				}
				if len(usage.Matches) == 0 {
					if _, err = red.Fprintln(w, "      - MISSING DEFINITION!"); err != nil {
						return err
					}
				} else {
					for _, match := range usage.Matches {
						if _, err = fmt.Fprintf(w, "      - %s - %s (%s)\n", match.Index.Name, match.Field.Type, match.Field.Kind); err != nil {
							return err
						}
						if match.Position.IsValid() {
							if _, err = fmt.Fprintf(w, "        %s:%d:%d\n", match.Position.Filename, match.Position.Line, match.Position.Column); err != nil {
								return err
							}
						}
					}
				}
			}
		}

	}

	return nil
}
