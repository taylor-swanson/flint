// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package kql

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractKqlFields(t *testing.T) {
	tests := []struct {
		name   string
		query  string
		fields []string
	}{
		{
			name:   "simple field query",
			query:  `host.name: "web-01"`,
			fields: []string{"host.name"},
		},
		{
			name:   "range query",
			query:  `http.response.status_code >= 500`,
			fields: []string{"http.response.status_code"},
		},
		{
			name:   "exists query",
			query:  `event.category: *`,
			fields: []string{"event.category"},
		},
		{
			name:   "boolean and",
			query:  `host.name: "web-01" AND http.response.status_code: 500`,
			fields: []string{"host.name", "http.response.status_code"},
		},
		{
			name:   "boolean or",
			query:  `event.category: "process" OR event.category: "network"`,
			fields: []string{"event.category"},
		},
		{
			name:   "not query",
			query:  `NOT user.name: "SYSTEM"`,
			fields: []string{"user.name"},
		},
		{
			name:   "quoted field name",
			query:  `"@timestamp": "2024-01-01"`,
			fields: []string{"@timestamp"},
		},
		{
			name:   "wildcard field pattern",
			query:  `event.*: "process"`,
			fields: []string{"event.*"},
		},
		{
			name:   "nested query",
			query:  `process: { name: "cmd.exe" }`,
			fields: []string{"process.name"},
		},
		{
			name:   "deeply nested query",
			query:  `threat: { enrichments: { indicator: { type: "domain" } } }`,
			fields: []string{"threat.enrichments.indicator.type"},
		},
		{
			name:   "nested query with multiple inner fields",
			query:  `process: { name: "cmd.exe" AND pid: 4 }`,
			fields: []string{"process.name", "process.pid"},
		},
		{
			name:   "fieldless query returns no fields",
			query:  `"some search text"`,
			fields: nil,
		},
		{
			name:   "match all returns no fields",
			query:  `*`,
			fields: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractFields(tt.query)
			require.NoError(t, err)
			if len(tt.fields) == 0 {
				assert.Empty(t, got)
			} else {
				assert.Equal(t, tt.fields, got)
			}
		})
	}
}
