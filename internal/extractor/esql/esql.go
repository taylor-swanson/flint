// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package esql

import (
	"sort"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/taylor-swanson/flint/internal/lang/esql"
)

// extractor walks an ES|QL parse tree and collects field references.
//
// Fields appear in three distinct places in the grammar:
//   - qualifiedName in #dereference (value expressions)
//   - qualifiedName in matchBooleanExpression (field:value syntax)
//   - qualifiedName in mvExpand, changePoint, fuse, etc.
//   - qualifiedNamePattern in KEEP, DROP, RENAME, ENRICH
//
// Rather than enumerate every command rule, we listen on EnterQualifiedName
// and EnterQualifiedNamePattern, which covers all of the above.
//
// Note: qualifiedName also appears as the assignment target in EVAL and STATS
// (e.g., `EVAL alias = expr`). Those are output field definitions, but they are
// field names in the query all the same and are included here.
type extractor struct {
	*esql.BaseEsqlBaseParserListener
	seen    map[string]struct{}
	indices []string
}

func (e *extractor) add(field string) {
	e.seen[field] = struct{}{}
}

func (e *extractor) EnterFromCommand(ctx *esql.FromCommandContext) {
	ipmf := ctx.IndexPatternAndMetadataFields()
	if ipmf == nil {
		return
	}
	for _, ipos := range ipmf.AllIndexPatternOrSubquery() {
		ip := ipos.IndexPattern()
		if ip == nil {
			continue
		}
		if uis := ip.UnquotedIndexString(); uis != nil {
			e.indices = append(e.indices, uis.GetText())
		} else if is := ip.IndexString(); is != nil {
			e.indices = append(e.indices, strings.Trim(is.GetText(), `"`))
		}
	}
}

// EnterQualifiedName captures field references in expressions, match predicates,
// and command arguments (MV_EXPAND, CHANGE_POINT, FUSE, etc.).
func (e *extractor) EnterQualifiedName(ctx *esql.QualifiedNameContext) {
	e.add(ctx.GetText())
}

// EnterQualifiedNamePattern captures field selectors in KEEP, DROP, RENAME,
// and ENRICH commands. These may contain wildcards (e.g. `host.*`).
func (e *extractor) EnterQualifiedNamePattern(ctx *esql.QualifiedNamePatternContext) {
	e.add(ctx.GetText())
}

// ExtractFields extracts the indices and fields from an ESQL query.
func ExtractFields(query string) ([]string, []string, error) {
	lexer := esql.NewEsqlBaseLexer(antlr.NewInputStream(query))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := esql.NewEsqlBaseParser(stream)

	tree := p.SingleStatement()

	e := &extractor{seen: make(map[string]struct{})}
	antlr.ParseTreeWalkerDefault.Walk(e, tree)

	sort.Strings(e.indices)

	var fields []string
	for k := range e.seen {
		fields = append(fields, k)
	}
	sort.Strings(fields)

	return e.indices, fields, nil
}
