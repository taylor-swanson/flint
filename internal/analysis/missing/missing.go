// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package missing

import (
	"fmt"

	"github.com/taylor-swanson/flint/internal/analysis"
	"github.com/taylor-swanson/flint/internal/ecs"
)

const Name = "missing"

var Analyzer = &analysis.Analyzer{
	Name:      Name,
	Doc:       "Detect missing definitions for field usages",
	RunUsages: runUsages,
}

const ecsRef = "v9.4.0"

func runUsages(pass *analysis.UsagePass) error {
	ecsResolver, err := ecs.NewResolver(ecsRef, pass.CacheDir)
	if err != nil {
		return err
	}

	for _, u := range pass.Usages {
		if len(u.Matches) > 0 {
			continue
		}

		// ECS fields can be resolved by the dynamic template. If this is an
		// ECS field, then don't error on a missing definition.
		if f := ecsResolver.Lookup(u.Name); f != nil {
			continue
		}

		issue := analysis.Issue{
			Analyzer: Name,
			Message:  fmt.Sprintf("Missing definition for %q with index pattern %q", u.Name, u.Pattern),
			Severity: analysis.SeverityError,
			Pos:      u.Position,
		}

		pass.Issues = append(pass.Issues, issue)
	}

	return nil
}
