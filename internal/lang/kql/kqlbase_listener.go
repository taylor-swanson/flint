// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// Code generated from KqlBase.g4 by ANTLR 4.13.2. DO NOT EDIT.

package kql // KqlBase
import "github.com/antlr4-go/antlr/v4"

// KqlBaseListener is a complete listener for a parse tree produced by KqlBaseParser.
type KqlBaseListener interface {
	antlr.ParseTreeListener

	// EnterTopLevelQuery is called when entering the topLevelQuery production.
	EnterTopLevelQuery(c *TopLevelQueryContext)

	// EnterBooleanQuery is called when entering the booleanQuery production.
	EnterBooleanQuery(c *BooleanQueryContext)

	// EnterDefaultQuery is called when entering the defaultQuery production.
	EnterDefaultQuery(c *DefaultQueryContext)

	// EnterSimpleQuery is called when entering the simpleQuery production.
	EnterSimpleQuery(c *SimpleQueryContext)

	// EnterNotQuery is called when entering the notQuery production.
	EnterNotQuery(c *NotQueryContext)

	// EnterNestedQuery is called when entering the nestedQuery production.
	EnterNestedQuery(c *NestedQueryContext)

	// EnterBooleanNestedQuery is called when entering the booleanNestedQuery production.
	EnterBooleanNestedQuery(c *BooleanNestedQueryContext)

	// EnterDefaultNestedQuery is called when entering the defaultNestedQuery production.
	EnterDefaultNestedQuery(c *DefaultNestedQueryContext)

	// EnterNestedSimpleSubQuery is called when entering the nestedSimpleSubQuery production.
	EnterNestedSimpleSubQuery(c *NestedSimpleSubQueryContext)

	// EnterNestedParenthesizedQuery is called when entering the nestedParenthesizedQuery production.
	EnterNestedParenthesizedQuery(c *NestedParenthesizedQueryContext)

	// EnterMatchAllQuery is called when entering the matchAllQuery production.
	EnterMatchAllQuery(c *MatchAllQueryContext)

	// EnterParenthesizedQuery is called when entering the parenthesizedQuery production.
	EnterParenthesizedQuery(c *ParenthesizedQueryContext)

	// EnterRangeQuery is called when entering the rangeQuery production.
	EnterRangeQuery(c *RangeQueryContext)

	// EnterRangeQueryValue is called when entering the rangeQueryValue production.
	EnterRangeQueryValue(c *RangeQueryValueContext)

	// EnterExistsQuery is called when entering the existsQuery production.
	EnterExistsQuery(c *ExistsQueryContext)

	// EnterFieldQuery is called when entering the fieldQuery production.
	EnterFieldQuery(c *FieldQueryContext)

	// EnterFieldLessQuery is called when entering the fieldLessQuery production.
	EnterFieldLessQuery(c *FieldLessQueryContext)

	// EnterFieldQueryValue is called when entering the fieldQueryValue production.
	EnterFieldQueryValue(c *FieldQueryValueContext)

	// EnterFieldQueryValueLiteral is called when entering the fieldQueryValueLiteral production.
	EnterFieldQueryValueLiteral(c *FieldQueryValueLiteralContext)

	// EnterFieldQueryValueUnquotedLiteral is called when entering the fieldQueryValueUnquotedLiteral production.
	EnterFieldQueryValueUnquotedLiteral(c *FieldQueryValueUnquotedLiteralContext)

	// EnterBooleanFieldQueryValue is called when entering the booleanFieldQueryValue production.
	EnterBooleanFieldQueryValue(c *BooleanFieldQueryValueContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// ExitTopLevelQuery is called when exiting the topLevelQuery production.
	ExitTopLevelQuery(c *TopLevelQueryContext)

	// ExitBooleanQuery is called when exiting the booleanQuery production.
	ExitBooleanQuery(c *BooleanQueryContext)

	// ExitDefaultQuery is called when exiting the defaultQuery production.
	ExitDefaultQuery(c *DefaultQueryContext)

	// ExitSimpleQuery is called when exiting the simpleQuery production.
	ExitSimpleQuery(c *SimpleQueryContext)

	// ExitNotQuery is called when exiting the notQuery production.
	ExitNotQuery(c *NotQueryContext)

	// ExitNestedQuery is called when exiting the nestedQuery production.
	ExitNestedQuery(c *NestedQueryContext)

	// ExitBooleanNestedQuery is called when exiting the booleanNestedQuery production.
	ExitBooleanNestedQuery(c *BooleanNestedQueryContext)

	// ExitDefaultNestedQuery is called when exiting the defaultNestedQuery production.
	ExitDefaultNestedQuery(c *DefaultNestedQueryContext)

	// ExitNestedSimpleSubQuery is called when exiting the nestedSimpleSubQuery production.
	ExitNestedSimpleSubQuery(c *NestedSimpleSubQueryContext)

	// ExitNestedParenthesizedQuery is called when exiting the nestedParenthesizedQuery production.
	ExitNestedParenthesizedQuery(c *NestedParenthesizedQueryContext)

	// ExitMatchAllQuery is called when exiting the matchAllQuery production.
	ExitMatchAllQuery(c *MatchAllQueryContext)

	// ExitParenthesizedQuery is called when exiting the parenthesizedQuery production.
	ExitParenthesizedQuery(c *ParenthesizedQueryContext)

	// ExitRangeQuery is called when exiting the rangeQuery production.
	ExitRangeQuery(c *RangeQueryContext)

	// ExitRangeQueryValue is called when exiting the rangeQueryValue production.
	ExitRangeQueryValue(c *RangeQueryValueContext)

	// ExitExistsQuery is called when exiting the existsQuery production.
	ExitExistsQuery(c *ExistsQueryContext)

	// ExitFieldQuery is called when exiting the fieldQuery production.
	ExitFieldQuery(c *FieldQueryContext)

	// ExitFieldLessQuery is called when exiting the fieldLessQuery production.
	ExitFieldLessQuery(c *FieldLessQueryContext)

	// ExitFieldQueryValue is called when exiting the fieldQueryValue production.
	ExitFieldQueryValue(c *FieldQueryValueContext)

	// ExitFieldQueryValueLiteral is called when exiting the fieldQueryValueLiteral production.
	ExitFieldQueryValueLiteral(c *FieldQueryValueLiteralContext)

	// ExitFieldQueryValueUnquotedLiteral is called when exiting the fieldQueryValueUnquotedLiteral production.
	ExitFieldQueryValueUnquotedLiteral(c *FieldQueryValueUnquotedLiteralContext)

	// ExitBooleanFieldQueryValue is called when exiting the booleanFieldQueryValue production.
	ExitBooleanFieldQueryValue(c *BooleanFieldQueryValueContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)
}
