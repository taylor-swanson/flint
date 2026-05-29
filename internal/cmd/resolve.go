// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/taylor-swanson/flint/internal/consumer"
	"github.com/taylor-swanson/flint/internal/field"
	"github.com/taylor-swanson/flint/internal/flint"
	"github.com/taylor-swanson/flint/internal/provider"
)

func newCmdResolve() *cobra.Command {
	var providerConfigs kvListFlag
	var usageType string
	var indexPattern string
	var outputJSON bool
	var minifyJSON bool
	var noColor bool

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

			report := flint.NewReport(usages)

			if outputJSON {
				err = flint.PrintJSON(os.Stdout, &report, minifyJSON)
			} else {
				err = flint.PrintText(os.Stdout, &report, noColor)
			}
			if err != nil {
				return fmt.Errorf("failed to print report: %w", err)
			}

			return nil
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
