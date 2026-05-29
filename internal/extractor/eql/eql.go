// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package eql

import (
	"sort"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/taylor-swanson/flint/internal/lang/eql"
)

// extractor walks an EQL parse tree and collects field references.
// Field references are qualifiedName nodes in the #dereference alternative
// of primaryExpression — the only place fields appear in the grammar.
type extractor struct {
	*eql.BaseEqlBaseListener
	seen map[string]struct{}
}

func (e *extractor) add(field string) {
	e.seen[field] = struct{}{}
}

// EnterDereference is called for every qualifiedName used as a value expression.
// These are the only field references in EQL — function names, event categories,
// pipe names, and relationship keywords all live in separate grammar rules.
func (e *extractor) EnterDereference(ctx *eql.DereferenceContext) {
	name := ctx.QualifiedName().GetText()
	// Strip the optional-access marker prefix
	name = strings.TrimPrefix(name, "?")
	e.add(name)
}

// ExtractFields extracts fields from an EQL query.
func ExtractFields(query string) ([]string, error) {
	lexer := eql.NewEqlBaseLexer(antlr.NewInputStream(query))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := eql.NewEqlBaseParser(stream)

	tree := p.SingleStatement()

	e := &extractor{seen: make(map[string]struct{})}
	antlr.ParseTreeWalkerDefault.Walk(e, tree)

	var fields []string
	for k := range e.seen {
		fields = append(fields, k)
	}
	sort.Strings(fields)

	return fields, nil
}
