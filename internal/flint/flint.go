// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package flint

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/taylor-swanson/flint/internal/consumer"
	"github.com/taylor-swanson/flint/internal/field"
	"github.com/taylor-swanson/flint/internal/provider"
	"github.com/taylor-swanson/flint/internal/version"
)

var (
	providerConfigs kvListFlag
	usageType       string
	indexPattern    string
	cacheDir        string
	paths           []string
	outputJSON      bool
	minifyJSON      bool
	noColor         bool
	showVersion     bool
)

func parseArgs() {
	flag.Var(&providerConfigs, "provider", "field definition providers")
	flag.StringVar(&usageType, "type", "", "usage type")
	flag.StringVar(&indexPattern, "index-pattern", "", "index pattern to use when resolving definitions (if applicable)")
	flag.StringVar(&cacheDir, "cache", ".cache", "cache directory")
	flag.BoolVar(&outputJSON, "json", false, "output as JSON")
	flag.BoolVar(&minifyJSON, "minify", false, "minify json output")
	flag.BoolVar(&noColor, "no-color", false, "disable color output")
	flag.BoolVar(&showVersion, "version", false, "show version")

	flag.Usage = func() {
		out := flag.CommandLine.Output()

		_, _ = fmt.Fprintf(out, "%s [flags] PATH [PATH ...]\n", version.Name)
		_, _ = fmt.Fprintln(out)

		_, _ = fmt.Fprintln(out, "Args:")
		_, _ = fmt.Fprintln(out, "  PATH    path to scan for usages (provide multiple times for additional paths)")
		_, _ = fmt.Fprintln(out)

		_, _ = fmt.Fprintln(out, "Flags:")
		flag.PrintDefaults()
	}
	flag.Parse()

	paths = flag.Args()
}

func Run() int {
	parseArgs()

	if showVersion {
		fmt.Printf("%s version %s [commit %v]\n", version.Name, version.Version, version.Commit)
		return 0
	}

	if len(providerConfigs) == 0 {
		out := flag.CommandLine.Output()
		_, _ = fmt.Fprintln(out, "At least one provider is required (specify via -provider flag)")
		_, _ = fmt.Fprintln(out)
		flag.Usage()
		return 2
	}
	if usageType == "" {
		out := flag.CommandLine.Output()
		_, _ = fmt.Fprintln(out, "A usage type is required (specify via -type flag)")
		_, _ = fmt.Fprintln(out)
		flag.Usage()
		return 2
	}
	c, err := consumer.New(usageType, indexPattern, paths)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return 2
	}

	if len(paths) == 0 {
		out := flag.CommandLine.Output()
		_, _ = fmt.Fprintln(out, "At least one path is required (specify via PATH arg)")
		_, _ = fmt.Fprintln(out)
		flag.Usage()
		return 2
	}

	var providers []provider.Provider
	for _, pc := range providerConfigs {
		p, err := provider.New(pc.Key, pc.Value)
		if err != nil {
			slog.Error("Failed to create provider", slog.String("error", err.Error()))
			os.Exit(1)
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
			slog.Error("Failed to provide fields", slog.String("error", err.Error()))
			return 1
		}
	}

	usages, err := c.Consume(&consumer.State{
		Ctx:      context.Background(),
		FieldMap: fieldMap,
		CacheDir: cacheDir,
	})
	if err != nil {
		slog.Error("Failed to get usages", slog.String("error", err.Error()))
		return 1
	}

	report := NewReport(usages)

	if outputJSON {
		err = PrintJSON(os.Stdout, &report, minifyJSON)
	} else {
		err = PrintText(os.Stdout, &report, noColor)
	}

	if err != nil {
		slog.Error("Failed to print report", slog.String("error", err.Error()))
		return 1
	}

	return 0
}
