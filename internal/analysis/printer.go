// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package analysis

import (
	"encoding/json"
	"fmt"
	"io"
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

func PrintText(w io.Writer, report *Report, wantColor bool) error {
	red := color.New(color.FgRed)
	yellow := color.New(color.FgYellow)
	green := color.New(color.FgGreen)
	bold := color.New(color.Bold)

	if !wantColor {
		red.DisableColor()
		yellow.DisableColor()
		green.DisableColor()
		bold.DisableColor()
	}

	var err error
	if _, err = bold.Fprint(w, "Time:"); err != nil {
		return err
	}
	if _, err = fmt.Fprintf(w, " %s\n\n", report.Timestamp.Format(time.RFC3339)); err != nil {
		return err
	}

	if _, err = bold.Fprintln(w, "Files:"); err != nil {
		return err
	}
	for _, name := range report.Files {
		if _, err = fmt.Fprintf(w, "  - %s\n", name); err != nil {
			return err
		}
	}
	if _, err = fmt.Fprintln(w); err != nil {
		return err
	}

	if _, err = bold.Fprintln(w, "Analyzers:"); err != nil {
		return err
	}
	for _, name := range report.Analyzers {
		if _, err = fmt.Fprintf(w, "  - %s\n", name); err != nil {
			return err
		}
	}
	if _, err = fmt.Fprintln(w); err != nil {
		return err
	}

	// -------------------------------------------------------------------------
	// File Reports

	for _, filename := range report.Files {

		fr := report.FileReports[filename]
		if len(fr.UsageReports) == 0 {
			continue
		}

		if _, err = bold.Fprintln(w, "Usage reports for:"); err != nil {
			return err
		}
		if _, err = bold.Fprintln(w, filename); err != nil {
			return err
		}

		fmt.Println()

		var total int
		for _, analyzerName := range report.Analyzers {
			usageReport, ok := fr.UsageReports[analyzerName]
			if !ok || len(usageReport.Issues) == 0 {
				continue
			}
			total += len(usageReport.Issues)

			if _, err = bold.Fprintln(w, analyzerName); err != nil {
				return err
			}

			for _, issue := range usageReport.Issues {
				if _, err = fmt.Fprint(w, "  - "); err != nil {
					return err
				}

				switch issue.Severity {
				case SeverityHint:
					if _, err = fmt.Fprint(w, "HINT: "+issue.Message); err != nil {
						return err
					}
				case SeverityInfo:
					if _, err = fmt.Fprint(w, "INFO: "+issue.Message); err != nil {
						return err
					}
				case SeverityWarn:
					if _, err = yellow.Fprint(w, "WARN: "+issue.Message); err != nil {
						return err
					}
				case SeverityError:
					if _, err = red.Fprint(w, "ERROR: "+issue.Message); err != nil {
						return err
					}
				case SeverityDeprecated:
					if _, err = yellow.Fprint(w, "DEPRECATED: "+issue.Message); err != nil {
						return err
					}
				}

				if _, err = fmt.Fprintf(w, " (%s)\n", issue.Analyzer); err != nil {
					return err
				}

				if issue.Pos.IsValid() {
					if _, err = bold.Fprintf(w, "    %s\n", issue.Pos); err != nil {
						return err
					}
				}
				if len(issue.Related) > 0 {
					for _, related := range issue.Related {
						if _, err = fmt.Fprintf(w, "      - %s\n", related.Message); err != nil {
							return err
						}

						if related.Pos.IsValid() {
							if _, err = bold.Fprintf(w, "        %s\n", related.Pos); err != nil {
								return err
							}
						}
					}
				}
			}

			if _, err = fmt.Fprintln(w); err != nil {
				return err
			}
		}

		// Issue totals

		if total > 0 {
			if _, err = red.Fprintf(w, "%d issues:\n", total); err != nil {
				return err
			}
		} else {
			if _, err = green.Fprintln(w, "No issues"); err != nil {
				return err
			}
		}
		for _, analyzerName := range report.Analyzers {
			usageReport, ok := fr.UsageReports[analyzerName]
			if !ok || len(usageReport.Issues) == 0 {
				continue
			}
			if _, err = fmt.Fprintf(w, "  - %s: %d\n", analyzerName, len(usageReport.Issues)); err != nil {
				return err
			}
		}
		if _, err = fmt.Fprintln(w); err != nil {
			return err
		}
	}

	return nil
}
