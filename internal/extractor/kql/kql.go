// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package kql

import (
	"sort"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/taylor-swanson/flint/internal/lang/kql"
)

// extractor walks a KQL parse tree and collects field references as
// fully-qualified dot-separated names.
//
// In KQL, every field reference is a fieldName node. It appears in:
//   - fieldQuery:  fieldName COLON fieldQueryValue
//   - rangeQuery:  fieldName operator rangeQueryValue
//   - existsQuery: fieldName COLON WILDCARD
//   - nestedQuery: fieldName COLON LEFT_CURLY_BRACKET nestedSubQuery RIGHT_CURLY_BRACKET
//
// For nested queries (e.g. `process: { name: "cmd.exe" }`), the parent field
// name is pushed onto a stack so that inner field names are emitted as
// fully-qualified paths (e.g. "process.name"). Arbitrary nesting depth is
// supported.
type extractor struct {
	*kql.BaseKqlBaseListener
	seen        map[string]struct{}
	nestedStack []string
}

func (e *extractor) add(field string) {
	if field == "" {
		return
	}
	e.seen[field] = struct{}{}
}

// EnterNestedQuery pushes the parent field name onto the nesting stack.
// Inner field names will be prefixed with this path.
func (e *extractor) EnterNestedQuery(ctx *kql.NestedQueryContext) {
	name := stripQuotes(ctx.FieldName().GetValue().GetText())
	e.nestedStack = append(e.nestedStack, name)
}

// ExitNestedQuery pops the parent field name from the nesting stack.
func (e *extractor) ExitNestedQuery(_ *kql.NestedQueryContext) {
	if len(e.nestedStack) > 0 {
		e.nestedStack = e.nestedStack[:len(e.nestedStack)-1]
	}
}

// EnterFieldName is called for every fieldName node in the tree. Nodes that
// are the key of a nestedQuery are skipped here — they are handled in
// EnterNestedQuery to build the nesting stack instead.
func (e *extractor) EnterFieldName(ctx *kql.FieldNameContext) {
	if _, ok := ctx.GetParent().(*kql.NestedQueryContext); ok {
		return
	}
	name := stripQuotes(ctx.GetValue().GetText())
	if len(e.nestedStack) > 0 {
		name = strings.Join(append(e.nestedStack, name), ".")
	}
	e.add(name)
}

func stripQuotes(s string) string {
	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		return s[1 : len(s)-1]
	}
	return s
}

// ExtractFields extracts fields from a KQL query.
func ExtractFields(query string) ([]string, error) {
	lexer := kql.NewKqlBaseLexer(antlr.NewInputStream(query))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := kql.NewKqlBaseParser(stream)

	tree := p.TopLevelQuery()

	e := &extractor{seen: make(map[string]struct{})}
	antlr.ParseTreeWalkerDefault.Walk(e, tree)

	var fields []string
	for k := range e.seen {
		fields = append(fields, k)
	}
	sort.Strings(fields)

	return fields, nil
}
