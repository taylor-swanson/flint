// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package eql

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractFields(t *testing.T) {
	tests := []struct {
		name   string
		query  string
		fields []string
	}{
		{
			name:   "simple process query",
			query:  `process where process.name == "cmd.exe" and process.pid != 4`,
			fields: []string{"process.name", "process.pid"},
		},
		{
			name:   "network query with in predicate",
			query:  `network where source.ip == "10.0.0.1" and destination.port in (80, 443)`,
			fields: []string{"destination.port", "source.ip"},
		},
		{
			name:   "wildcard string match and negation",
			query:  `file where file.path : "C:\\Windows\\*" and not user.name == "SYSTEM"`,
			fields: []string{"file.path", "user.name"},
		},
		{
			name: "sequence with shared join key",
			query: `sequence by host.id
			  [process where process.name == "powershell.exe"]
			  [network where destination.ip != null]`,
			fields: []string{"destination.ip", "host.id", "process.name"},
		},
		{
			name:   "optional field access",
			query:  `process where ?process.parent.name == "explorer.exe"`,
			fields: []string{"process.parent.name"},
		},
		{
			name:   "function call arguments are still fields",
			query:  `process where match(process.command_line, ".*powershell.*")`,
			fields: []string{"process.command_line"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractFields(tt.query)
			require.NoError(t, err)
			assert.Equal(t, tt.fields, got)
		})
	}
}
