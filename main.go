// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// The flint command provides a linter and lookup tool for fields.
package main

import (
	"log/slog"
	"os"

	"github.com/taylor-swanson/flint/internal/cmd"

	_ "github.com/taylor-swanson/flint/internal/consumer/detectionrules"
	_ "github.com/taylor-swanson/flint/internal/consumer/generic"
	_ "github.com/taylor-swanson/flint/internal/consumer/integrations"

	_ "github.com/taylor-swanson/flint/internal/provider/beats"
	_ "github.com/taylor-swanson/flint/internal/provider/integrations"
)

func main() {
	if err := cmd.Execute(); err != nil {
		slog.Error("Error running app", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
