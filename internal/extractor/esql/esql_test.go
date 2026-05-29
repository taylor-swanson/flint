// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package esql

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractEsqlFields(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		indices []string
		fields  []string
	}{
		{
			name:    "simple where filter",
			query:   `FROM logs | WHERE host.name == "web-01" AND http.status_code >= 500`,
			indices: []string{"logs"},
			fields:  []string{"host.name", "http.status_code"},
		},
		{
			name:    "eval with alias",
			query:   `FROM logs | EVAL upper_host = TO_UPPER(host.name)`,
			indices: []string{"logs"},
			fields:  []string{"host.name", "upper_host"},
		},
		{
			name:    "stats by group",
			query:   `FROM logs | STATS count = COUNT(*) BY service.name`,
			indices: []string{"logs"},
			fields:  []string{"count", "service.name"},
		},
		{
			name:    "keep command uses qualifiedNamePattern",
			query:   `FROM logs | KEEP host.name, http.*`,
			indices: []string{"logs"},
			fields:  []string{"host.name", "http.*"},
		},
		{
			name:    "drop command",
			query:   `FROM logs | DROP event.original`,
			indices: []string{"logs"},
			fields:  []string{"event.original"},
		},
		{
			name:    "rename command",
			query:   `FROM logs | RENAME host.name AS hostname`,
			indices: []string{"logs"},
			fields:  []string{"host.name", "hostname"},
		},
		{
			name:    "match predicate (field:value syntax)",
			query:   `FROM logs | WHERE message:"error" AND log.level:"warn"`,
			indices: []string{"logs"},
			fields:  []string{"log.level", "message"},
		},
		{
			name:    "mv_expand",
			query:   `FROM logs | MV_EXPAND tags`,
			indices: []string{"logs"},
			fields:  []string{"tags"},
		},
		{
			name:    "sort",
			query:   `FROM logs | SORT @timestamp DESC`,
			indices: []string{"logs"},
			fields:  []string{"@timestamp"},
		},
		{
			name:    "pipeline with multiple commands",
			query:   `FROM logs | WHERE event.category == "process" | STATS c = COUNT(*) BY host.name | SORT c DESC`,
			indices: []string{"logs"},
			fields:  []string{"c", "event.category", "host.name"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndices, gotFields, err := ExtractFields(tt.query)
			require.NoError(t, err)
			assert.Equal(t, tt.indices, gotIndices)
			assert.Equal(t, tt.fields, gotFields)
		})
	}
}
