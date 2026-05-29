// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package detectionrules

import (
	"fmt"
	"sort"

	"github.com/BurntSushi/toml"

	"github.com/taylor-swanson/flint/internal/extractor/eql"
	"github.com/taylor-swanson/flint/internal/extractor/esql"
	"github.com/taylor-swanson/flint/internal/extractor/kql"
)

type rule struct {
	Rule ruleData `toml:"rule"`
}

type ruleData struct {
	RuleID                string               `toml:"rule_id"`
	Query                 string               `toml:"query,omitempty"`
	Language              string               `toml:"language,omitempty"`
	TimestampOverride     string               `toml:"timestamp_override,omitempty"`
	TimestampField        string               `toml:"timestamp_field,omitempty"`
	TiebreakerField       string               `toml:"tiebreaker_field,omitempty"`
	EventCategoryOverride string               `toml:"event_category_override,omitempty"`
	RuleNameOverride      string               `toml:"rule_name_override,omitempty"`
	InvestigationFields   *investigationFields `toml:"investigation_fields,omitempty"`
	RequiredFields        []requiredField      `toml:"required_fields,omitempty"`
	AlertSuppression      *alertSuppression    `toml:"alert_suppression,omitempty"`
	RiskScoreMapping      []riskScoreMapping   `toml:"risk_score_mapping,omitempty"`
	SeverityMapping       []severityMapping    `toml:"severity_mapping,omitempty"`
	Threshold             *thresholdMapping    `toml:"threshold,omitempty"`
	NewTerms              *newTermsMapping     `toml:"new_terms,omitempty"`
	ThreatMapping         []threatMapEntries   `toml:"threat_mapping,omitempty"`
	Index                 []string             `toml:"index,omitempty"`
}

type investigationFields struct {
	FieldNames []string `toml:"field_names"`
}

type requiredField struct {
	Name string `toml:"name"`
}

type alertSuppression struct {
	GroupBy []string `toml:"group_by"`
}

type riskScoreMapping struct {
	Field string `toml:"field"`
}

type severityMapping struct {
	Field string `toml:"field"`
}

type thresholdMapping struct {
	Field       []string               `toml:"field"`
	Cardinality []thresholdCardinality `toml:"cardinality,omitempty"`
}

type thresholdCardinality struct {
	Field string `toml:"field"`
}

type newTermsMapping struct {
	Field              string               `toml:"field"`
	HistoryWindowStart []historyWindowStart `toml:"history_window_start"`
}

type historyWindowStart struct {
	Field string `toml:"field"`
}

type threatMapEntries struct {
	Entries []threatMapEntry `toml:"entries"`
}

type threatMapEntry struct {
	Field string `toml:"field"`
}

// parse extracts indices and fields from a detection rule.
func parse(data []byte) ([]string, []string, error) {
	var r rule
	var err error

	if err = toml.Unmarshal(data, &r); err != nil {
		return nil, nil, err
	}

	found := map[string]struct{}{}

	var indices []string
	var queryFields []string
	switch r.Rule.Language {
	case "eql":
		queryFields, err = eql.ExtractFields(r.Rule.Query)
	case "esql":
		indices, queryFields, err = esql.ExtractFields(r.Rule.Query)
	case "kuery", "lucene":
		queryFields, err = kql.ExtractFields(r.Rule.Query)
	default:
		return nil, nil, fmt.Errorf("unknown rule language: %q", r.Rule.Language)
	}
	if err != nil {
		return nil, nil, fmt.Errorf("failed to extract query fields: %w", err)
	}

	for _, field := range queryFields {
		found[field] = struct{}{}
	}

	found[r.Rule.TimestampOverride] = struct{}{}
	found[r.Rule.TimestampField] = struct{}{}

	if r.Rule.InvestigationFields != nil {
		for _, v := range r.Rule.InvestigationFields.FieldNames {
			found[v] = struct{}{}
		}
	}

	for _, f := range r.Rule.RequiredFields {
		found[f.Name] = struct{}{}
	}

	if r.Rule.AlertSuppression != nil {
		for _, v := range r.Rule.AlertSuppression.GroupBy {
			found[v] = struct{}{}
		}
	}

	for _, m := range r.Rule.RiskScoreMapping {
		found[m.Field] = struct{}{}
	}

	for _, m := range r.Rule.SeverityMapping {
		found[m.Field] = struct{}{}
	}

	if r.Rule.Threshold != nil {
		for _, v := range r.Rule.Threshold.Field {
			found[v] = struct{}{}
		}
		for _, v := range r.Rule.Threshold.Cardinality {
			found[v.Field] = struct{}{}
		}
	}

	if r.Rule.NewTerms != nil {
		found[r.Rule.NewTerms.Field] = struct{}{}
		for _, v := range r.Rule.NewTerms.HistoryWindowStart {
			found[v.Field] = struct{}{}
		}
	}

	for _, entries := range r.Rule.ThreatMapping {
		for _, e := range entries.Entries {
			found[e.Field] = struct{}{}
		}
	}

	indices = append(indices, r.Rule.Index...)

	sort.Strings(indices)

	fields := make([]string, 0, len(found))
	for field := range found {
		if field == "" {
			continue
		}
		fields = append(fields, field)
	}
	sort.Strings(fields)

	return indices, fields, nil
}
