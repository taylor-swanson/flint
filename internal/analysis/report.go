// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package analysis

import "time"

type IssueReport struct {
	Issues []Issue `json:"issues,omitempty"`
}

type FileReport struct {
	UsageReports map[string]IssueReport `json:"usage_reports,omitempty"`
}

type Report struct {
	Timestamp   time.Time             `json:"timestamp"`
	Files       []string              `json:"files,omitempty"`
	Analyzers   []string              `json:"analyzers,omitempty"`
	FileReports map[string]FileReport `json:"file_reports,omitempty"`
}
