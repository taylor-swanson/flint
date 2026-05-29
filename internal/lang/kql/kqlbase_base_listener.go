// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// Code generated from KqlBase.g4 by ANTLR 4.13.2. DO NOT EDIT.

package kql // KqlBase
import "github.com/antlr4-go/antlr/v4"

// BaseKqlBaseListener is a complete listener for a parse tree produced by KqlBaseParser.
type BaseKqlBaseListener struct{}

var _ KqlBaseListener = &BaseKqlBaseListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseKqlBaseListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseKqlBaseListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseKqlBaseListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseKqlBaseListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTopLevelQuery is called when production topLevelQuery is entered.
func (s *BaseKqlBaseListener) EnterTopLevelQuery(ctx *TopLevelQueryContext) {}

// ExitTopLevelQuery is called when production topLevelQuery is exited.
func (s *BaseKqlBaseListener) ExitTopLevelQuery(ctx *TopLevelQueryContext) {}

// EnterBooleanQuery is called when production booleanQuery is entered.
func (s *BaseKqlBaseListener) EnterBooleanQuery(ctx *BooleanQueryContext) {}

// ExitBooleanQuery is called when production booleanQuery is exited.
func (s *BaseKqlBaseListener) ExitBooleanQuery(ctx *BooleanQueryContext) {}

// EnterDefaultQuery is called when production defaultQuery is entered.
func (s *BaseKqlBaseListener) EnterDefaultQuery(ctx *DefaultQueryContext) {}

// ExitDefaultQuery is called when production defaultQuery is exited.
func (s *BaseKqlBaseListener) ExitDefaultQuery(ctx *DefaultQueryContext) {}

// EnterSimpleQuery is called when production simpleQuery is entered.
func (s *BaseKqlBaseListener) EnterSimpleQuery(ctx *SimpleQueryContext) {}

// ExitSimpleQuery is called when production simpleQuery is exited.
func (s *BaseKqlBaseListener) ExitSimpleQuery(ctx *SimpleQueryContext) {}

// EnterNotQuery is called when production notQuery is entered.
func (s *BaseKqlBaseListener) EnterNotQuery(ctx *NotQueryContext) {}

// ExitNotQuery is called when production notQuery is exited.
func (s *BaseKqlBaseListener) ExitNotQuery(ctx *NotQueryContext) {}

// EnterNestedQuery is called when production nestedQuery is entered.
func (s *BaseKqlBaseListener) EnterNestedQuery(ctx *NestedQueryContext) {}

// ExitNestedQuery is called when production nestedQuery is exited.
func (s *BaseKqlBaseListener) ExitNestedQuery(ctx *NestedQueryContext) {}

// EnterBooleanNestedQuery is called when production booleanNestedQuery is entered.
func (s *BaseKqlBaseListener) EnterBooleanNestedQuery(ctx *BooleanNestedQueryContext) {}

// ExitBooleanNestedQuery is called when production booleanNestedQuery is exited.
func (s *BaseKqlBaseListener) ExitBooleanNestedQuery(ctx *BooleanNestedQueryContext) {}

// EnterDefaultNestedQuery is called when production defaultNestedQuery is entered.
func (s *BaseKqlBaseListener) EnterDefaultNestedQuery(ctx *DefaultNestedQueryContext) {}

// ExitDefaultNestedQuery is called when production defaultNestedQuery is exited.
func (s *BaseKqlBaseListener) ExitDefaultNestedQuery(ctx *DefaultNestedQueryContext) {}

// EnterNestedSimpleSubQuery is called when production nestedSimpleSubQuery is entered.
func (s *BaseKqlBaseListener) EnterNestedSimpleSubQuery(ctx *NestedSimpleSubQueryContext) {}

// ExitNestedSimpleSubQuery is called when production nestedSimpleSubQuery is exited.
func (s *BaseKqlBaseListener) ExitNestedSimpleSubQuery(ctx *NestedSimpleSubQueryContext) {}

// EnterNestedParenthesizedQuery is called when production nestedParenthesizedQuery is entered.
func (s *BaseKqlBaseListener) EnterNestedParenthesizedQuery(ctx *NestedParenthesizedQueryContext) {}

// ExitNestedParenthesizedQuery is called when production nestedParenthesizedQuery is exited.
func (s *BaseKqlBaseListener) ExitNestedParenthesizedQuery(ctx *NestedParenthesizedQueryContext) {}

// EnterMatchAllQuery is called when production matchAllQuery is entered.
func (s *BaseKqlBaseListener) EnterMatchAllQuery(ctx *MatchAllQueryContext) {}

// ExitMatchAllQuery is called when production matchAllQuery is exited.
func (s *BaseKqlBaseListener) ExitMatchAllQuery(ctx *MatchAllQueryContext) {}

// EnterParenthesizedQuery is called when production parenthesizedQuery is entered.
func (s *BaseKqlBaseListener) EnterParenthesizedQuery(ctx *ParenthesizedQueryContext) {}

// ExitParenthesizedQuery is called when production parenthesizedQuery is exited.
func (s *BaseKqlBaseListener) ExitParenthesizedQuery(ctx *ParenthesizedQueryContext) {}

// EnterRangeQuery is called when production rangeQuery is entered.
func (s *BaseKqlBaseListener) EnterRangeQuery(ctx *RangeQueryContext) {}

// ExitRangeQuery is called when production rangeQuery is exited.
func (s *BaseKqlBaseListener) ExitRangeQuery(ctx *RangeQueryContext) {}

// EnterRangeQueryValue is called when production rangeQueryValue is entered.
func (s *BaseKqlBaseListener) EnterRangeQueryValue(ctx *RangeQueryValueContext) {}

// ExitRangeQueryValue is called when production rangeQueryValue is exited.
func (s *BaseKqlBaseListener) ExitRangeQueryValue(ctx *RangeQueryValueContext) {}

// EnterExistsQuery is called when production existsQuery is entered.
func (s *BaseKqlBaseListener) EnterExistsQuery(ctx *ExistsQueryContext) {}

// ExitExistsQuery is called when production existsQuery is exited.
func (s *BaseKqlBaseListener) ExitExistsQuery(ctx *ExistsQueryContext) {}

// EnterFieldQuery is called when production fieldQuery is entered.
func (s *BaseKqlBaseListener) EnterFieldQuery(ctx *FieldQueryContext) {}

// ExitFieldQuery is called when production fieldQuery is exited.
func (s *BaseKqlBaseListener) ExitFieldQuery(ctx *FieldQueryContext) {}

// EnterFieldLessQuery is called when production fieldLessQuery is entered.
func (s *BaseKqlBaseListener) EnterFieldLessQuery(ctx *FieldLessQueryContext) {}

// ExitFieldLessQuery is called when production fieldLessQuery is exited.
func (s *BaseKqlBaseListener) ExitFieldLessQuery(ctx *FieldLessQueryContext) {}

// EnterFieldQueryValue is called when production fieldQueryValue is entered.
func (s *BaseKqlBaseListener) EnterFieldQueryValue(ctx *FieldQueryValueContext) {}

// ExitFieldQueryValue is called when production fieldQueryValue is exited.
func (s *BaseKqlBaseListener) ExitFieldQueryValue(ctx *FieldQueryValueContext) {}

// EnterFieldQueryValueLiteral is called when production fieldQueryValueLiteral is entered.
func (s *BaseKqlBaseListener) EnterFieldQueryValueLiteral(ctx *FieldQueryValueLiteralContext) {}

// ExitFieldQueryValueLiteral is called when production fieldQueryValueLiteral is exited.
func (s *BaseKqlBaseListener) ExitFieldQueryValueLiteral(ctx *FieldQueryValueLiteralContext) {}

// EnterFieldQueryValueUnquotedLiteral is called when production fieldQueryValueUnquotedLiteral is entered.
func (s *BaseKqlBaseListener) EnterFieldQueryValueUnquotedLiteral(ctx *FieldQueryValueUnquotedLiteralContext) {
}

// ExitFieldQueryValueUnquotedLiteral is called when production fieldQueryValueUnquotedLiteral is exited.
func (s *BaseKqlBaseListener) ExitFieldQueryValueUnquotedLiteral(ctx *FieldQueryValueUnquotedLiteralContext) {
}

// EnterBooleanFieldQueryValue is called when production booleanFieldQueryValue is entered.
func (s *BaseKqlBaseListener) EnterBooleanFieldQueryValue(ctx *BooleanFieldQueryValueContext) {}

// ExitBooleanFieldQueryValue is called when production booleanFieldQueryValue is exited.
func (s *BaseKqlBaseListener) ExitBooleanFieldQueryValue(ctx *BooleanFieldQueryValueContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseKqlBaseListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseKqlBaseListener) ExitFieldName(ctx *FieldNameContext) {}
