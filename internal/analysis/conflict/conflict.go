package conflict

import (
	"fmt"

	"github.com/taylor-swanson/flint/internal/analysis"
	"github.com/taylor-swanson/flint/internal/field"
)

const Name = "conflict"

var Analyzer = &analysis.Analyzer{
	Name:      Name,
	Doc:       "Check for conflicting declarations for field usages",
	RunUsages: runUsages,
}

func runUsages(pass *analysis.UsagePass) error {
	//seen := map[string]map[string][]*field.Field{}
	seen := map[string]map[string]map[*field.Field][]*field.Definition{}

	for _, u := range pass.Usages {
		m, ok := seen[u.Pattern]
		if !ok {
			m = map[string]map[*field.Field][]*field.Definition{}
			seen[u.Pattern] = m
		}

		for _, def := range u.Matches {
			fieldUsages, ok := m[u.Name]
			if !ok {
				fieldUsages = map[*field.Field][]*field.Definition{}
				m[u.Name] = fieldUsages
			}

			fieldUsages[def.Field] = append(fieldUsages[def.Field], def)
		}
	}

	for indexPattern, fields := range seen {
		for fieldName, seenFields := range fields {
			if len(seenFields) < 2 {
				continue
			}

			issue := analysis.Issue{
				Analyzer: Name,
				Message:  fmt.Sprintf("Usage of field %q with index pattern %q has %d conflicting declarations", fieldName, indexPattern, len(seenFields)),
				Severity: analysis.SeverityError,
			}
			for f, defs := range seenFields {
				for _, def := range defs {
					issue.Related = append(issue.Related, analysis.Related{
						Message: fmt.Sprintf("Observed type %q in index %q", string(f.Type), def.Index.Name),
						Pos:     def.Position,
					})
				}
			}

			pass.Issues = append(pass.Issues, issue)
		}
	}

	return nil
}
