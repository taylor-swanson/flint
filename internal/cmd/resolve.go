// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/taylor-swanson/flint/internal/analysis"
	"github.com/taylor-swanson/flint/internal/analysis/conflict"
	"github.com/taylor-swanson/flint/internal/analysis/definition"
	"github.com/taylor-swanson/flint/internal/analysis/missing"
	"github.com/taylor-swanson/flint/internal/consumer"
	"github.com/taylor-swanson/flint/internal/field"
	"github.com/taylor-swanson/flint/internal/provider"
)

func newCmdResolve() *cobra.Command {
	var providerConfigs kvListFlag
	var usageType string
	var indexPattern string
	var outputJSON bool
	var minifyJSON bool
	var noColor bool

	analyzers := []*analysis.Analyzer{
		conflict.Analyzer,
		definition.Analyzer,
		missing.Analyzer,
	}
	analyzerNames := make([]string, len(analyzers))
	for i, a := range analyzers {
		analyzerNames[i] = a.Name
	}

	cmd := &cobra.Command{
		Use:               "resolve PATH [PATH ...]",
		Short:             "Resolve field definitions",
		Args:              cobra.MinimumNArgs(1),
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cacheDir, err := cmd.Flags().GetString("cache")
			if err != nil {
				return err
			}

			if len(providerConfigs) == 0 {
				return errors.New("at least one provider configuration is required")
			}
			if usageType == "" {
				return errors.New("a usage type is required (specify via --type flag)")
			}

			c, err := consumer.New(usageType, indexPattern, args)
			if err != nil {
				return fmt.Errorf("could not create consumer for %s: %w", usageType, err)
			}

			var providers []provider.Provider
			for _, pc := range providerConfigs {
				p, err := provider.New(pc.Key, pc.Value)
				if err != nil {
					return fmt.Errorf("could not create provider for %s: %w", usageType, err)
				}
				providers = append(providers, p)
			}

			fieldMap := field.NewMap()
			providerState := provider.State{
				Ctx:      context.Background(),
				FieldMap: fieldMap,
				CacheDir: cacheDir,
			}

			for _, p := range providers {
				if err = p.Provide(&providerState); err != nil {
					return fmt.Errorf("could not provide fields: %w", err)
				}
			}

			usages, err := c.Consume(&consumer.State{
				Ctx:      context.Background(),
				FieldMap: fieldMap,
				CacheDir: cacheDir,
			})
			if err != nil {
				return fmt.Errorf("failed to get usages: %w", err)
			}

			usagesByFile := make(map[string][]field.Usage, len(args))
			for _, u := range usages {
				usagesByFile[u.Name] = append(usagesByFile[u.Name], u)
			}

			report := analysis.Report{
				Timestamp:   time.Now(),
				Analyzers:   analyzerNames,
				Files:       args,
				FileReports: map[string]analysis.FileReport{},
			}

			for _, filename := range report.Files {
				fr := analysis.FileReport{UsageReports: map[string]analysis.IssueReport{}}

				for _, a := range analyzers {
					pass := analysis.UsagePass{
						Ctx:      context.Background(),
						Usages:   usages,
						CacheDir: cacheDir,
					}

					if a.RunUsages != nil {
						if err = a.RunUsages(&pass); err != nil {
							return err
						}
						fr.UsageReports[a.Name] = analysis.IssueReport{
							Issues: pass.Issues,
						}
					}
				}

				if len(fr.UsageReports) == 0 {
					continue
				}
				report.FileReports[filename] = fr
			}

			if outputJSON {
				err = analysis.PrintJSON(os.Stdout, &report, minifyJSON)
			} else {
				err = analysis.PrintText(os.Stdout, &report, !noColor)
			}

			return err
		},
	}

	cmd.Flags().VarP(&providerConfigs, "provider", "p", "field definition providers")
	cmd.Flags().StringVarP(&usageType, "type", "t", "", "field consumer type")
	cmd.Flags().StringVarP(&indexPattern, "index-pattern", "i", "", "index pattern to use when resolving definitions (if applicable)")
	cmd.Flags().BoolVarP(&outputJSON, "json", "j", false, "output as JSON")
	cmd.Flags().BoolVarP(&minifyJSON, "minify", "m", false, "minify JSON")
	cmd.Flags().BoolVar(&noColor, "no-color", false, "disable color output")

	return cmd
}
