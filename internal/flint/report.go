// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package flint

import (
	"sort"
	"time"

	"github.com/taylor-swanson/flint/internal/field"
)

type IndexReport struct {
	Usages []field.Usage `json:"usages"`
}

type FileReport struct {
	IndexReports map[string]*IndexReport `json:"index_reports"`
	Indices      []string                `json:"-"`
}

type Report struct {
	Timestamp   time.Time              `json:"timestamp"`
	FileReports map[string]*FileReport `json:"file_reports"`
	Files       []string               `json:"-"`
}

func NewReport(usages []field.Usage) Report {
	r := Report{
		Timestamp:   time.Now(),
		FileReports: map[string]*FileReport{},
	}

	for _, v := range usages {
		filename := v.Position.Filename

		fr, ok := r.FileReports[filename]
		if !ok {
			r.Files = append(r.Files, filename)
			fr = &FileReport{IndexReports: map[string]*IndexReport{}}
			r.FileReports[filename] = fr
		}

		ir, ok := fr.IndexReports[v.Pattern]
		if !ok {
			fr.Indices = append(fr.Indices, v.Pattern)
			ir = &IndexReport{}
			fr.IndexReports[v.Pattern] = ir
		}

		ir.Usages = append(ir.Usages, v)
	}

	sort.Strings(r.Files)
	for _, v := range r.FileReports {
		sort.Strings(v.Indices)
		for _, vv := range v.IndexReports {
			sort.Slice(vv.Usages, func(i, j int) bool {
				return vv.Usages[i].Name < vv.Usages[j].Name
			})
		}
	}

	return r
}
