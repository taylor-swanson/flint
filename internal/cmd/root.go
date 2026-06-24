// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/taylor-swanson/flint/internal/version"
)

func Execute() error {
	cmd := cobra.Command{
		Use:               version.Name,
		Short:             "Field linter and lookup tool",
		SilenceErrors:     true,
		DisableAutoGenTag: true,
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			debug, _ := cmd.Flags().GetBool("debug")

			level := slog.LevelInfo
			if debug {
				level = slog.LevelDebug
			}
			slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level})))
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}

	cmd.PersistentFlags().String("cache", ".cache", "path to cache directory")
	cmd.PersistentFlags().Bool("debug", false, "enable debug logging")

	cmd.AddCommand(
		newCmdResolve(),
		newCmdVersion(),
	)

	return cmd.Execute()
}
