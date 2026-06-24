package definition

import (
	"fmt"

	"github.com/taylor-swanson/flint/internal/analysis"
)

const Name = "definition"

var Analyzer = &analysis.Analyzer{
	Name:      Name,
	Doc:       "Show definitions for field usages",
	RunUsages: runUsages,
}

func runUsages(pass *analysis.UsagePass) error {
	for _, u := range pass.Usages {
		if len(u.Matches) == 0 {
			issue := analysis.Issue{
				Analyzer: Name,
				Message:  fmt.Sprintf("No definitions found for %q with index pattern %q", u.Name, u.Pattern),
				Severity: analysis.SeverityWarn,
				Pos:      u.Position,
			}
			pass.Issues = append(pass.Issues, issue)
			continue
		}

		issue := analysis.Issue{
			Analyzer: Name,
			Message:  fmt.Sprintf("Found %d definition(s) for %q with index pattern %q", len(u.Matches), u.Name, u.Pattern),
			Severity: analysis.SeverityInfo,
			Pos:      u.Position,
		}

		for _, def := range u.Matches {
			issue.Related = append(issue.Related, analysis.Related{
				Pos:     def.Position,
				Message: fmt.Sprintf("Defined in %s-%s", def.Index.Name, def.Index.Version),
			})
		}
		pass.Issues = append(pass.Issues, issue)
	}

	return nil
}
