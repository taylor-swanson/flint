// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// Code generated from KqlBase.g4 by ANTLR 4.13.2. DO NOT EDIT.

package kql // KqlBase
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

/*
 * Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
 * or more contributor license agreements. Licensed under the Elastic License
 * 2.0; you may not use this file except in compliance with the Elastic License
 * 2.0.
 */

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type KqlBaseParser struct {
	*antlr.BaseParser
}

var KqlBaseParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func kqlbaseParserInit() {
	staticData := &KqlBaseParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'and'", "'or'", "'not'", "':'", "'<'", "'<='", "'>'", "'>='",
		"'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "DEFAULT_SKIP", "AND", "OR", "NOT", "COLON", "OP_LESS", "OP_LESS_EQ",
		"OP_MORE", "OP_MORE_EQ", "LEFT_PARENTHESIS", "RIGHT_PARENTHESIS", "LEFT_CURLY_BRACKET",
		"RIGHT_CURLY_BRACKET", "UNQUOTED_LITERAL", "QUOTED_STRING", "WILDCARD",
	}
	staticData.RuleNames = []string{
		"topLevelQuery", "query", "simpleQuery", "notQuery", "nestedQuery",
		"nestedSubQuery", "nestedSimpleSubQuery", "nestedParenthesizedQuery",
		"matchAllQuery", "parenthesizedQuery", "rangeQuery", "rangeQueryValue",
		"existsQuery", "fieldQuery", "fieldLessQuery", "fieldQueryValue", "fieldQueryValueLiteral",
		"fieldQueryValueUnquotedLiteral", "booleanFieldQueryValue", "fieldName",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 16, 196, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 3, 0, 42,
		8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 52, 8, 1, 10,
		1, 12, 1, 55, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 65, 8, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 82, 8, 5, 10, 5, 12, 5, 85, 9, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 94, 8, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 3, 8, 102, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 4, 11, 115, 8, 11, 11, 11, 12, 11, 116,
		1, 11, 3, 11, 120, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		3, 15, 139, 8, 15, 1, 16, 1, 16, 3, 16, 143, 8, 16, 1, 17, 4, 17, 146,
		8, 17, 11, 17, 12, 17, 147, 1, 17, 5, 17, 151, 8, 17, 10, 17, 12, 17, 154,
		9, 17, 1, 17, 3, 17, 157, 8, 17, 1, 17, 1, 17, 4, 17, 161, 8, 17, 11, 17,
		12, 17, 162, 1, 17, 3, 17, 166, 8, 17, 1, 17, 1, 17, 3, 17, 170, 8, 17,
		1, 17, 3, 17, 173, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3,
		18, 181, 8, 18, 1, 18, 1, 18, 1, 18, 5, 18, 186, 8, 18, 10, 18, 12, 18,
		189, 9, 18, 1, 19, 1, 19, 1, 19, 3, 19, 194, 8, 19, 1, 19, 0, 3, 2, 10,
		36, 20, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 0, 4, 1, 0, 2, 3, 1, 0, 6, 9, 2, 0, 14, 14, 16, 16, 3, 0, 4,
		4, 14, 14, 16, 16, 211, 0, 41, 1, 0, 0, 0, 2, 45, 1, 0, 0, 0, 4, 64, 1,
		0, 0, 0, 6, 66, 1, 0, 0, 0, 8, 69, 1, 0, 0, 0, 10, 75, 1, 0, 0, 0, 12,
		93, 1, 0, 0, 0, 14, 95, 1, 0, 0, 0, 16, 101, 1, 0, 0, 0, 18, 105, 1, 0,
		0, 0, 20, 109, 1, 0, 0, 0, 22, 119, 1, 0, 0, 0, 24, 121, 1, 0, 0, 0, 26,
		125, 1, 0, 0, 0, 28, 129, 1, 0, 0, 0, 30, 138, 1, 0, 0, 0, 32, 142, 1,
		0, 0, 0, 34, 172, 1, 0, 0, 0, 36, 180, 1, 0, 0, 0, 38, 193, 1, 0, 0, 0,
		40, 42, 3, 2, 1, 0, 41, 40, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 43, 1,
		0, 0, 0, 43, 44, 5, 0, 0, 1, 44, 1, 1, 0, 0, 0, 45, 46, 6, 1, -1, 0, 46,
		47, 3, 4, 2, 0, 47, 53, 1, 0, 0, 0, 48, 49, 10, 2, 0, 0, 49, 50, 7, 0,
		0, 0, 50, 52, 3, 2, 1, 2, 51, 48, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51,
		1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 3, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0,
		56, 65, 3, 6, 3, 0, 57, 65, 3, 8, 4, 0, 58, 65, 3, 18, 9, 0, 59, 65, 3,
		16, 8, 0, 60, 65, 3, 24, 12, 0, 61, 65, 3, 20, 10, 0, 62, 65, 3, 26, 13,
		0, 63, 65, 3, 28, 14, 0, 64, 56, 1, 0, 0, 0, 64, 57, 1, 0, 0, 0, 64, 58,
		1, 0, 0, 0, 64, 59, 1, 0, 0, 0, 64, 60, 1, 0, 0, 0, 64, 61, 1, 0, 0, 0,
		64, 62, 1, 0, 0, 0, 64, 63, 1, 0, 0, 0, 65, 5, 1, 0, 0, 0, 66, 67, 5, 4,
		0, 0, 67, 68, 3, 4, 2, 0, 68, 7, 1, 0, 0, 0, 69, 70, 3, 38, 19, 0, 70,
		71, 5, 5, 0, 0, 71, 72, 5, 12, 0, 0, 72, 73, 3, 10, 5, 0, 73, 74, 5, 13,
		0, 0, 74, 9, 1, 0, 0, 0, 75, 76, 6, 5, -1, 0, 76, 77, 3, 12, 6, 0, 77,
		83, 1, 0, 0, 0, 78, 79, 10, 2, 0, 0, 79, 80, 7, 0, 0, 0, 80, 82, 3, 10,
		5, 2, 81, 78, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84,
		1, 0, 0, 0, 84, 11, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 94, 3, 6, 3, 0,
		87, 94, 3, 8, 4, 0, 88, 94, 3, 16, 8, 0, 89, 94, 3, 14, 7, 0, 90, 94, 3,
		24, 12, 0, 91, 94, 3, 20, 10, 0, 92, 94, 3, 26, 13, 0, 93, 86, 1, 0, 0,
		0, 93, 87, 1, 0, 0, 0, 93, 88, 1, 0, 0, 0, 93, 89, 1, 0, 0, 0, 93, 90,
		1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 13, 1, 0, 0, 0,
		95, 96, 5, 10, 0, 0, 96, 97, 3, 10, 5, 0, 97, 98, 5, 11, 0, 0, 98, 15,
		1, 0, 0, 0, 99, 100, 5, 16, 0, 0, 100, 102, 5, 5, 0, 0, 101, 99, 1, 0,
		0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 5, 16, 0, 0,
		104, 17, 1, 0, 0, 0, 105, 106, 5, 10, 0, 0, 106, 107, 3, 2, 1, 0, 107,
		108, 5, 11, 0, 0, 108, 19, 1, 0, 0, 0, 109, 110, 3, 38, 19, 0, 110, 111,
		7, 1, 0, 0, 111, 112, 3, 22, 11, 0, 112, 21, 1, 0, 0, 0, 113, 115, 7, 2,
		0, 0, 114, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0,
		116, 117, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 120, 5, 15, 0, 0, 119,
		114, 1, 0, 0, 0, 119, 118, 1, 0, 0, 0, 120, 23, 1, 0, 0, 0, 121, 122, 3,
		38, 19, 0, 122, 123, 5, 5, 0, 0, 123, 124, 5, 16, 0, 0, 124, 25, 1, 0,
		0, 0, 125, 126, 3, 38, 19, 0, 126, 127, 5, 5, 0, 0, 127, 128, 3, 30, 15,
		0, 128, 27, 1, 0, 0, 0, 129, 130, 3, 30, 15, 0, 130, 29, 1, 0, 0, 0, 131,
		139, 3, 32, 16, 0, 132, 133, 5, 4, 0, 0, 133, 139, 3, 30, 15, 0, 134, 135,
		5, 10, 0, 0, 135, 136, 3, 36, 18, 0, 136, 137, 5, 11, 0, 0, 137, 139, 1,
		0, 0, 0, 138, 131, 1, 0, 0, 0, 138, 132, 1, 0, 0, 0, 138, 134, 1, 0, 0,
		0, 139, 31, 1, 0, 0, 0, 140, 143, 3, 34, 17, 0, 141, 143, 5, 15, 0, 0,
		142, 140, 1, 0, 0, 0, 142, 141, 1, 0, 0, 0, 143, 33, 1, 0, 0, 0, 144, 146,
		7, 2, 0, 0, 145, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 145, 1, 0,
		0, 0, 147, 148, 1, 0, 0, 0, 148, 152, 1, 0, 0, 0, 149, 151, 7, 3, 0, 0,
		150, 149, 1, 0, 0, 0, 151, 154, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152,
		153, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 155, 157,
		7, 0, 0, 0, 156, 155, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 173, 1, 0,
		0, 0, 158, 169, 7, 0, 0, 0, 159, 161, 7, 3, 0, 0, 160, 159, 1, 0, 0, 0,
		161, 162, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163,
		165, 1, 0, 0, 0, 164, 166, 7, 0, 0, 0, 165, 164, 1, 0, 0, 0, 165, 166,
		1, 0, 0, 0, 166, 170, 1, 0, 0, 0, 167, 170, 5, 3, 0, 0, 168, 170, 5, 2,
		0, 0, 169, 160, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 168, 1, 0, 0, 0,
		169, 170, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 173, 5, 4, 0, 0, 172,
		145, 1, 0, 0, 0, 172, 158, 1, 0, 0, 0, 172, 171, 1, 0, 0, 0, 173, 35, 1,
		0, 0, 0, 174, 175, 6, 18, -1, 0, 175, 176, 5, 10, 0, 0, 176, 177, 3, 36,
		18, 0, 177, 178, 5, 11, 0, 0, 178, 181, 1, 0, 0, 0, 179, 181, 3, 30, 15,
		0, 180, 174, 1, 0, 0, 0, 180, 179, 1, 0, 0, 0, 181, 187, 1, 0, 0, 0, 182,
		183, 10, 3, 0, 0, 183, 184, 7, 0, 0, 0, 184, 186, 3, 30, 15, 0, 185, 182,
		1, 0, 0, 0, 186, 189, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 188, 1, 0,
		0, 0, 188, 37, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 190, 194, 5, 14, 0, 0,
		191, 194, 5, 15, 0, 0, 192, 194, 5, 16, 0, 0, 193, 190, 1, 0, 0, 0, 193,
		191, 1, 0, 0, 0, 193, 192, 1, 0, 0, 0, 194, 39, 1, 0, 0, 0, 20, 41, 53,
		64, 83, 93, 101, 116, 119, 138, 142, 147, 152, 156, 162, 165, 169, 172,
		180, 187, 193,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// KqlBaseParserInit initializes any static state used to implement KqlBaseParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewKqlBaseParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func KqlBaseParserInit() {
	staticData := &KqlBaseParserStaticData
	staticData.once.Do(kqlbaseParserInit)
}

// NewKqlBaseParser produces a new parser instance for the optional input antlr.TokenStream.
func NewKqlBaseParser(input antlr.TokenStream) *KqlBaseParser {
	KqlBaseParserInit()
	this := new(KqlBaseParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &KqlBaseParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "KqlBase.g4"

	return this
}

// KqlBaseParser tokens.
const (
	KqlBaseParserEOF                 = antlr.TokenEOF
	KqlBaseParserDEFAULT_SKIP        = 1
	KqlBaseParserAND                 = 2
	KqlBaseParserOR                  = 3
	KqlBaseParserNOT                 = 4
	KqlBaseParserCOLON               = 5
	KqlBaseParserOP_LESS             = 6
	KqlBaseParserOP_LESS_EQ          = 7
	KqlBaseParserOP_MORE             = 8
	KqlBaseParserOP_MORE_EQ          = 9
	KqlBaseParserLEFT_PARENTHESIS    = 10
	KqlBaseParserRIGHT_PARENTHESIS   = 11
	KqlBaseParserLEFT_CURLY_BRACKET  = 12
	KqlBaseParserRIGHT_CURLY_BRACKET = 13
	KqlBaseParserUNQUOTED_LITERAL    = 14
	KqlBaseParserQUOTED_STRING       = 15
	KqlBaseParserWILDCARD            = 16
)

// KqlBaseParser rules.
const (
	KqlBaseParserRULE_topLevelQuery                  = 0
	KqlBaseParserRULE_query                          = 1
	KqlBaseParserRULE_simpleQuery                    = 2
	KqlBaseParserRULE_notQuery                       = 3
	KqlBaseParserRULE_nestedQuery                    = 4
	KqlBaseParserRULE_nestedSubQuery                 = 5
	KqlBaseParserRULE_nestedSimpleSubQuery           = 6
	KqlBaseParserRULE_nestedParenthesizedQuery       = 7
	KqlBaseParserRULE_matchAllQuery                  = 8
	KqlBaseParserRULE_parenthesizedQuery             = 9
	KqlBaseParserRULE_rangeQuery                     = 10
	KqlBaseParserRULE_rangeQueryValue                = 11
	KqlBaseParserRULE_existsQuery                    = 12
	KqlBaseParserRULE_fieldQuery                     = 13
	KqlBaseParserRULE_fieldLessQuery                 = 14
	KqlBaseParserRULE_fieldQueryValue                = 15
	KqlBaseParserRULE_fieldQueryValueLiteral         = 16
	KqlBaseParserRULE_fieldQueryValueUnquotedLiteral = 17
	KqlBaseParserRULE_booleanFieldQueryValue         = 18
	KqlBaseParserRULE_fieldName                      = 19
)

// ITopLevelQueryContext is an interface to support dynamic dispatch.
type ITopLevelQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	Query() IQueryContext

	// IsTopLevelQueryContext differentiates from other interfaces.
	IsTopLevelQueryContext()
}

type TopLevelQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelQueryContext() *TopLevelQueryContext {
	var p = new(TopLevelQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_topLevelQuery
	return p
}

func InitEmptyTopLevelQueryContext(p *TopLevelQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_topLevelQuery
}

func (*TopLevelQueryContext) IsTopLevelQueryContext() {}

func NewTopLevelQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelQueryContext {
	var p = new(TopLevelQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_topLevelQuery

	return p
}

func (s *TopLevelQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelQueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserEOF, 0)
}

func (s *TopLevelQueryContext) Query() IQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *TopLevelQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterTopLevelQuery(s)
	}
}

func (s *TopLevelQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitTopLevelQuery(s)
	}
}

func (p *KqlBaseParser) TopLevelQuery() (localctx ITopLevelQueryContext) {
	localctx = NewTopLevelQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, KqlBaseParserRULE_topLevelQuery)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&115740) != 0 {
		{
			p.SetState(40)
			p.query(0)
		}

	}
	{
		p.SetState(43)
		p.Match(KqlBaseParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) CopyAll(ctx *QueryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BooleanQueryContext struct {
	QueryContext
	operator antlr.Token
}

func NewBooleanQueryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanQueryContext {
	var p = new(BooleanQueryContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *BooleanQueryContext) GetOperator() antlr.Token { return s.operator }

func (s *BooleanQueryContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *BooleanQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanQueryContext) AllQuery() []IQueryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQueryContext); ok {
			len++
		}
	}

	tst := make([]IQueryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQueryContext); ok {
			tst[i] = t.(IQueryContext)
			i++
		}
	}

	return tst
}

func (s *BooleanQueryContext) Query(i int) IQueryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *BooleanQueryContext) AND() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserAND, 0)
}

func (s *BooleanQueryContext) OR() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserOR, 0)
}

func (s *BooleanQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterBooleanQuery(s)
	}
}

func (s *BooleanQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitBooleanQuery(s)
	}
}

type DefaultQueryContext struct {
	QueryContext
}

func NewDefaultQueryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultQueryContext {
	var p = new(DefaultQueryContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *DefaultQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultQueryContext) SimpleQuery() ISimpleQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleQueryContext)
}

func (s *DefaultQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterDefaultQuery(s)
	}
}

func (s *DefaultQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitDefaultQuery(s)
	}
}

func (p *KqlBaseParser) Query() (localctx IQueryContext) {
	return p.query(0)
}

func (p *KqlBaseParser) query(_p int) (localctx IQueryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewQueryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQueryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, KqlBaseParserRULE_query, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewDefaultQueryContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(46)
		p.SimpleQuery()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBooleanQueryContext(p, NewQueryContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, KqlBaseParserRULE_query)
			p.SetState(48)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(49)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*BooleanQueryContext).operator = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == KqlBaseParserAND || _la == KqlBaseParserOR) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*BooleanQueryContext).operator = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(50)
				p.query(2)
			}

		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleQueryContext is an interface to support dynamic dispatch.
type ISimpleQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NotQuery() INotQueryContext
	NestedQuery() INestedQueryContext
	ParenthesizedQuery() IParenthesizedQueryContext
	MatchAllQuery() IMatchAllQueryContext
	ExistsQuery() IExistsQueryContext
	RangeQuery() IRangeQueryContext
	FieldQuery() IFieldQueryContext
	FieldLessQuery() IFieldLessQueryContext

	// IsSimpleQueryContext differentiates from other interfaces.
	IsSimpleQueryContext()
}

type SimpleQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleQueryContext() *SimpleQueryContext {
	var p = new(SimpleQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_simpleQuery
	return p
}

func InitEmptySimpleQueryContext(p *SimpleQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_simpleQuery
}

func (*SimpleQueryContext) IsSimpleQueryContext() {}

func NewSimpleQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleQueryContext {
	var p = new(SimpleQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_simpleQuery

	return p
}

func (s *SimpleQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleQueryContext) NotQuery() INotQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotQueryContext)
}

func (s *SimpleQueryContext) NestedQuery() INestedQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedQueryContext)
}

func (s *SimpleQueryContext) ParenthesizedQuery() IParenthesizedQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParenthesizedQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParenthesizedQueryContext)
}

func (s *SimpleQueryContext) MatchAllQuery() IMatchAllQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchAllQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchAllQueryContext)
}

func (s *SimpleQueryContext) ExistsQuery() IExistsQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExistsQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExistsQueryContext)
}

func (s *SimpleQueryContext) RangeQuery() IRangeQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeQueryContext)
}

func (s *SimpleQueryContext) FieldQuery() IFieldQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldQueryContext)
}

func (s *SimpleQueryContext) FieldLessQuery() IFieldLessQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldLessQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldLessQueryContext)
}

func (s *SimpleQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterSimpleQuery(s)
	}
}

func (s *SimpleQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitSimpleQuery(s)
	}
}

func (p *KqlBaseParser) SimpleQuery() (localctx ISimpleQueryContext) {
	localctx = NewSimpleQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, KqlBaseParserRULE_simpleQuery)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.NotQuery()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.NestedQuery()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(58)
			p.ParenthesizedQuery()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(59)
			p.MatchAllQuery()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(60)
			p.ExistsQuery()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(61)
			p.RangeQuery()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(62)
			p.FieldQuery()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(63)
			p.FieldLessQuery()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INotQueryContext is an interface to support dynamic dispatch.
type INotQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSubQuery returns the subQuery rule contexts.
	GetSubQuery() ISimpleQueryContext

	// SetSubQuery sets the subQuery rule contexts.
	SetSubQuery(ISimpleQueryContext)

	// Getter signatures
	NOT() antlr.TerminalNode
	SimpleQuery() ISimpleQueryContext

	// IsNotQueryContext differentiates from other interfaces.
	IsNotQueryContext()
}

type NotQueryContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	subQuery ISimpleQueryContext
}

func NewEmptyNotQueryContext() *NotQueryContext {
	var p = new(NotQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_notQuery
	return p
}

func InitEmptyNotQueryContext(p *NotQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_notQuery
}

func (*NotQueryContext) IsNotQueryContext() {}

func NewNotQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotQueryContext {
	var p = new(NotQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_notQuery

	return p
}

func (s *NotQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *NotQueryContext) GetSubQuery() ISimpleQueryContext { return s.subQuery }

func (s *NotQueryContext) SetSubQuery(v ISimpleQueryContext) { s.subQuery = v }

func (s *NotQueryContext) NOT() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserNOT, 0)
}

func (s *NotQueryContext) SimpleQuery() ISimpleQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleQueryContext)
}

func (s *NotQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterNotQuery(s)
	}
}

func (s *NotQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitNotQuery(s)
	}
}

func (p *KqlBaseParser) NotQuery() (localctx INotQueryContext) {
	localctx = NewNotQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, KqlBaseParserRULE_notQuery)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(KqlBaseParserNOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)

		var _x = p.SimpleQuery()

		localctx.(*NotQueryContext).subQuery = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INestedQueryContext is an interface to support dynamic dispatch.
type INestedQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldName() IFieldNameContext
	COLON() antlr.TerminalNode
	LEFT_CURLY_BRACKET() antlr.TerminalNode
	NestedSubQuery() INestedSubQueryContext
	RIGHT_CURLY_BRACKET() antlr.TerminalNode

	// IsNestedQueryContext differentiates from other interfaces.
	IsNestedQueryContext()
}

type NestedQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedQueryContext() *NestedQueryContext {
	var p = new(NestedQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_nestedQuery
	return p
}

func InitEmptyNestedQueryContext(p *NestedQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_nestedQuery
}

func (*NestedQueryContext) IsNestedQueryContext() {}

func NewNestedQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedQueryContext {
	var p = new(NestedQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_nestedQuery

	return p
}

func (s *NestedQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedQueryContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *NestedQueryContext) COLON() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserCOLON, 0)
}

func (s *NestedQueryContext) LEFT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserLEFT_CURLY_BRACKET, 0)
}

func (s *NestedQueryContext) NestedSubQuery() INestedSubQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedSubQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedSubQueryContext)
}

func (s *NestedQueryContext) RIGHT_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserRIGHT_CURLY_BRACKET, 0)
}

func (s *NestedQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestedQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterNestedQuery(s)
	}
}

func (s *NestedQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitNestedQuery(s)
	}
}

func (p *KqlBaseParser) NestedQuery() (localctx INestedQueryContext) {
	localctx = NewNestedQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, KqlBaseParserRULE_nestedQuery)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.FieldName()
	}
	{
		p.SetState(70)
		p.Match(KqlBaseParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(71)
		p.Match(KqlBaseParserLEFT_CURLY_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)
		p.nestedSubQuery(0)
	}
	{
		p.SetState(73)
		p.Match(KqlBaseParserRIGHT_CURLY_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INestedSubQueryContext is an interface to support dynamic dispatch.
type INestedSubQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsNestedSubQueryContext differentiates from other interfaces.
	IsNestedSubQueryContext()
}

type NestedSubQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedSubQueryContext() *NestedSubQueryContext {
	var p = new(NestedSubQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_nestedSubQuery
	return p
}

func InitEmptyNestedSubQueryContext(p *NestedSubQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_nestedSubQuery
}

func (*NestedSubQueryContext) IsNestedSubQueryContext() {}

func NewNestedSubQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedSubQueryContext {
	var p = new(NestedSubQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_nestedSubQuery

	return p
}

func (s *NestedSubQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedSubQueryContext) CopyAll(ctx *NestedSubQueryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *NestedSubQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedSubQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BooleanNestedQueryContext struct {
	NestedSubQueryContext
	operator antlr.Token
}

func NewBooleanNestedQueryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanNestedQueryContext {
	var p = new(BooleanNestedQueryContext)

	InitEmptyNestedSubQueryContext(&p.NestedSubQueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*NestedSubQueryContext))

	return p
}

func (s *BooleanNestedQueryContext) GetOperator() antlr.Token { return s.operator }

func (s *BooleanNestedQueryContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *BooleanNestedQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanNestedQueryContext) AllNestedSubQuery() []INestedSubQueryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INestedSubQueryContext); ok {
			len++
		}
	}

	tst := make([]INestedSubQueryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INestedSubQueryContext); ok {
			tst[i] = t.(INestedSubQueryContext)
			i++
		}
	}

	return tst
}

func (s *BooleanNestedQueryContext) NestedSubQuery(i int) INestedSubQueryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedSubQueryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedSubQueryContext)
}

func (s *BooleanNestedQueryContext) AND() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserAND, 0)
}

func (s *BooleanNestedQueryContext) OR() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserOR, 0)
}

func (s *BooleanNestedQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterBooleanNestedQuery(s)
	}
}

func (s *BooleanNestedQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitBooleanNestedQuery(s)
	}
}

type DefaultNestedQueryContext struct {
	NestedSubQueryContext
}

func NewDefaultNestedQueryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultNestedQueryContext {
	var p = new(DefaultNestedQueryContext)

	InitEmptyNestedSubQueryContext(&p.NestedSubQueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*NestedSubQueryContext))

	return p
}

func (s *DefaultNestedQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultNestedQueryContext) NestedSimpleSubQuery() INestedSimpleSubQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedSimpleSubQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedSimpleSubQueryContext)
}

func (s *DefaultNestedQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterDefaultNestedQuery(s)
	}
}

func (s *DefaultNestedQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitDefaultNestedQuery(s)
	}
}

func (p *KqlBaseParser) NestedSubQuery() (localctx INestedSubQueryContext) {
	return p.nestedSubQuery(0)
}

func (p *KqlBaseParser) nestedSubQuery(_p int) (localctx INestedSubQueryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewNestedSubQueryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx INestedSubQueryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, KqlBaseParserRULE_nestedSubQuery, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewDefaultNestedQueryContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(76)
		p.NestedSimpleSubQuery()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBooleanNestedQueryContext(p, NewNestedSubQueryContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, KqlBaseParserRULE_nestedSubQuery)
			p.SetState(78)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(79)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*BooleanNestedQueryContext).operator = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == KqlBaseParserAND || _la == KqlBaseParserOR) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*BooleanNestedQueryContext).operator = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(80)
				p.nestedSubQuery(2)
			}

		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INestedSimpleSubQueryContext is an interface to support dynamic dispatch.
type INestedSimpleSubQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NotQuery() INotQueryContext
	NestedQuery() INestedQueryContext
	MatchAllQuery() IMatchAllQueryContext
	NestedParenthesizedQuery() INestedParenthesizedQueryContext
	ExistsQuery() IExistsQueryContext
	RangeQuery() IRangeQueryContext
	FieldQuery() IFieldQueryContext

	// IsNestedSimpleSubQueryContext differentiates from other interfaces.
	IsNestedSimpleSubQueryContext()
}

type NestedSimpleSubQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedSimpleSubQueryContext() *NestedSimpleSubQueryContext {
	var p = new(NestedSimpleSubQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_nestedSimpleSubQuery
	return p
}

func InitEmptyNestedSimpleSubQueryContext(p *NestedSimpleSubQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_nestedSimpleSubQuery
}

func (*NestedSimpleSubQueryContext) IsNestedSimpleSubQueryContext() {}

func NewNestedSimpleSubQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedSimpleSubQueryContext {
	var p = new(NestedSimpleSubQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_nestedSimpleSubQuery

	return p
}

func (s *NestedSimpleSubQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedSimpleSubQueryContext) NotQuery() INotQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotQueryContext)
}

func (s *NestedSimpleSubQueryContext) NestedQuery() INestedQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedQueryContext)
}

func (s *NestedSimpleSubQueryContext) MatchAllQuery() IMatchAllQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchAllQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchAllQueryContext)
}

func (s *NestedSimpleSubQueryContext) NestedParenthesizedQuery() INestedParenthesizedQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedParenthesizedQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedParenthesizedQueryContext)
}

func (s *NestedSimpleSubQueryContext) ExistsQuery() IExistsQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExistsQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExistsQueryContext)
}

func (s *NestedSimpleSubQueryContext) RangeQuery() IRangeQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeQueryContext)
}

func (s *NestedSimpleSubQueryContext) FieldQuery() IFieldQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldQueryContext)
}

func (s *NestedSimpleSubQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedSimpleSubQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestedSimpleSubQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterNestedSimpleSubQuery(s)
	}
}

func (s *NestedSimpleSubQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitNestedSimpleSubQuery(s)
	}
}

func (p *KqlBaseParser) NestedSimpleSubQuery() (localctx INestedSimpleSubQueryContext) {
	localctx = NewNestedSimpleSubQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, KqlBaseParserRULE_nestedSimpleSubQuery)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.NotQuery()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.NestedQuery()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.MatchAllQuery()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(89)
			p.NestedParenthesizedQuery()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(90)
			p.ExistsQuery()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(91)
			p.RangeQuery()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(92)
			p.FieldQuery()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INestedParenthesizedQueryContext is an interface to support dynamic dispatch.
type INestedParenthesizedQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_PARENTHESIS() antlr.TerminalNode
	NestedSubQuery() INestedSubQueryContext
	RIGHT_PARENTHESIS() antlr.TerminalNode

	// IsNestedParenthesizedQueryContext differentiates from other interfaces.
	IsNestedParenthesizedQueryContext()
}

type NestedParenthesizedQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedParenthesizedQueryContext() *NestedParenthesizedQueryContext {
	var p = new(NestedParenthesizedQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_nestedParenthesizedQuery
	return p
}

func InitEmptyNestedParenthesizedQueryContext(p *NestedParenthesizedQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_nestedParenthesizedQuery
}

func (*NestedParenthesizedQueryContext) IsNestedParenthesizedQueryContext() {}

func NewNestedParenthesizedQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedParenthesizedQueryContext {
	var p = new(NestedParenthesizedQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_nestedParenthesizedQuery

	return p
}

func (s *NestedParenthesizedQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedParenthesizedQueryContext) LEFT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserLEFT_PARENTHESIS, 0)
}

func (s *NestedParenthesizedQueryContext) NestedSubQuery() INestedSubQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedSubQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedSubQueryContext)
}

func (s *NestedParenthesizedQueryContext) RIGHT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserRIGHT_PARENTHESIS, 0)
}

func (s *NestedParenthesizedQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedParenthesizedQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestedParenthesizedQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterNestedParenthesizedQuery(s)
	}
}

func (s *NestedParenthesizedQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitNestedParenthesizedQuery(s)
	}
}

func (p *KqlBaseParser) NestedParenthesizedQuery() (localctx INestedParenthesizedQueryContext) {
	localctx = NewNestedParenthesizedQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, KqlBaseParserRULE_nestedParenthesizedQuery)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(KqlBaseParserLEFT_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.nestedSubQuery(0)
	}
	{
		p.SetState(97)
		p.Match(KqlBaseParserRIGHT_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatchAllQueryContext is an interface to support dynamic dispatch.
type IMatchAllQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWILDCARD() []antlr.TerminalNode
	WILDCARD(i int) antlr.TerminalNode
	COLON() antlr.TerminalNode

	// IsMatchAllQueryContext differentiates from other interfaces.
	IsMatchAllQueryContext()
}

type MatchAllQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchAllQueryContext() *MatchAllQueryContext {
	var p = new(MatchAllQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_matchAllQuery
	return p
}

func InitEmptyMatchAllQueryContext(p *MatchAllQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_matchAllQuery
}

func (*MatchAllQueryContext) IsMatchAllQueryContext() {}

func NewMatchAllQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchAllQueryContext {
	var p = new(MatchAllQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_matchAllQuery

	return p
}

func (s *MatchAllQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchAllQueryContext) AllWILDCARD() []antlr.TerminalNode {
	return s.GetTokens(KqlBaseParserWILDCARD)
}

func (s *MatchAllQueryContext) WILDCARD(i int) antlr.TerminalNode {
	return s.GetToken(KqlBaseParserWILDCARD, i)
}

func (s *MatchAllQueryContext) COLON() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserCOLON, 0)
}

func (s *MatchAllQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchAllQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchAllQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterMatchAllQuery(s)
	}
}

func (s *MatchAllQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitMatchAllQuery(s)
	}
}

func (p *KqlBaseParser) MatchAllQuery() (localctx IMatchAllQueryContext) {
	localctx = NewMatchAllQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, KqlBaseParserRULE_matchAllQuery)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(99)
			p.Match(KqlBaseParserWILDCARD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Match(KqlBaseParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(103)
		p.Match(KqlBaseParserWILDCARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParenthesizedQueryContext is an interface to support dynamic dispatch.
type IParenthesizedQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_PARENTHESIS() antlr.TerminalNode
	Query() IQueryContext
	RIGHT_PARENTHESIS() antlr.TerminalNode

	// IsParenthesizedQueryContext differentiates from other interfaces.
	IsParenthesizedQueryContext()
}

type ParenthesizedQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParenthesizedQueryContext() *ParenthesizedQueryContext {
	var p = new(ParenthesizedQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_parenthesizedQuery
	return p
}

func InitEmptyParenthesizedQueryContext(p *ParenthesizedQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_parenthesizedQuery
}

func (*ParenthesizedQueryContext) IsParenthesizedQueryContext() {}

func NewParenthesizedQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParenthesizedQueryContext {
	var p = new(ParenthesizedQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_parenthesizedQuery

	return p
}

func (s *ParenthesizedQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *ParenthesizedQueryContext) LEFT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserLEFT_PARENTHESIS, 0)
}

func (s *ParenthesizedQueryContext) Query() IQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *ParenthesizedQueryContext) RIGHT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserRIGHT_PARENTHESIS, 0)
}

func (s *ParenthesizedQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesizedQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParenthesizedQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterParenthesizedQuery(s)
	}
}

func (s *ParenthesizedQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitParenthesizedQuery(s)
	}
}

func (p *KqlBaseParser) ParenthesizedQuery() (localctx IParenthesizedQueryContext) {
	localctx = NewParenthesizedQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, KqlBaseParserRULE_parenthesizedQuery)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(KqlBaseParserLEFT_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.query(0)
	}
	{
		p.SetState(107)
		p.Match(KqlBaseParserRIGHT_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeQueryContext is an interface to support dynamic dispatch.
type IRangeQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the operator token.
	GetOperator() antlr.Token

	// SetOperator sets the operator token.
	SetOperator(antlr.Token)

	// Getter signatures
	FieldName() IFieldNameContext
	RangeQueryValue() IRangeQueryValueContext
	OP_LESS() antlr.TerminalNode
	OP_LESS_EQ() antlr.TerminalNode
	OP_MORE() antlr.TerminalNode
	OP_MORE_EQ() antlr.TerminalNode

	// IsRangeQueryContext differentiates from other interfaces.
	IsRangeQueryContext()
}

type RangeQueryContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	operator antlr.Token
}

func NewEmptyRangeQueryContext() *RangeQueryContext {
	var p = new(RangeQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_rangeQuery
	return p
}

func InitEmptyRangeQueryContext(p *RangeQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_rangeQuery
}

func (*RangeQueryContext) IsRangeQueryContext() {}

func NewRangeQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeQueryContext {
	var p = new(RangeQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_rangeQuery

	return p
}

func (s *RangeQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeQueryContext) GetOperator() antlr.Token { return s.operator }

func (s *RangeQueryContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *RangeQueryContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *RangeQueryContext) RangeQueryValue() IRangeQueryValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeQueryValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeQueryValueContext)
}

func (s *RangeQueryContext) OP_LESS() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserOP_LESS, 0)
}

func (s *RangeQueryContext) OP_LESS_EQ() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserOP_LESS_EQ, 0)
}

func (s *RangeQueryContext) OP_MORE() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserOP_MORE, 0)
}

func (s *RangeQueryContext) OP_MORE_EQ() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserOP_MORE_EQ, 0)
}

func (s *RangeQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterRangeQuery(s)
	}
}

func (s *RangeQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitRangeQuery(s)
	}
}

func (p *KqlBaseParser) RangeQuery() (localctx IRangeQueryContext) {
	localctx = NewRangeQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, KqlBaseParserRULE_rangeQuery)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.FieldName()
	}
	{
		p.SetState(110)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*RangeQueryContext).operator = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&960) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*RangeQueryContext).operator = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(111)
		p.RangeQueryValue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeQueryValueContext is an interface to support dynamic dispatch.
type IRangeQueryValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUNQUOTED_LITERAL() []antlr.TerminalNode
	UNQUOTED_LITERAL(i int) antlr.TerminalNode
	AllWILDCARD() []antlr.TerminalNode
	WILDCARD(i int) antlr.TerminalNode
	QUOTED_STRING() antlr.TerminalNode

	// IsRangeQueryValueContext differentiates from other interfaces.
	IsRangeQueryValueContext()
}

type RangeQueryValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeQueryValueContext() *RangeQueryValueContext {
	var p = new(RangeQueryValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_rangeQueryValue
	return p
}

func InitEmptyRangeQueryValueContext(p *RangeQueryValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_rangeQueryValue
}

func (*RangeQueryValueContext) IsRangeQueryValueContext() {}

func NewRangeQueryValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeQueryValueContext {
	var p = new(RangeQueryValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_rangeQueryValue

	return p
}

func (s *RangeQueryValueContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeQueryValueContext) AllUNQUOTED_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(KqlBaseParserUNQUOTED_LITERAL)
}

func (s *RangeQueryValueContext) UNQUOTED_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(KqlBaseParserUNQUOTED_LITERAL, i)
}

func (s *RangeQueryValueContext) AllWILDCARD() []antlr.TerminalNode {
	return s.GetTokens(KqlBaseParserWILDCARD)
}

func (s *RangeQueryValueContext) WILDCARD(i int) antlr.TerminalNode {
	return s.GetToken(KqlBaseParserWILDCARD, i)
}

func (s *RangeQueryValueContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserQUOTED_STRING, 0)
}

func (s *RangeQueryValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeQueryValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeQueryValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterRangeQueryValue(s)
	}
}

func (s *RangeQueryValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitRangeQueryValue(s)
	}
}

func (p *KqlBaseParser) RangeQueryValue() (localctx IRangeQueryValueContext) {
	localctx = NewRangeQueryValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, KqlBaseParserRULE_rangeQueryValue)
	var _la int

	var _alt int

	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case KqlBaseParserUNQUOTED_LITERAL, KqlBaseParserWILDCARD:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(113)
					_la = p.GetTokenStream().LA(1)

					if !(_la == KqlBaseParserUNQUOTED_LITERAL || _la == KqlBaseParserWILDCARD) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(116)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case KqlBaseParserQUOTED_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Match(KqlBaseParserQUOTED_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExistsQueryContext is an interface to support dynamic dispatch.
type IExistsQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldName() IFieldNameContext
	COLON() antlr.TerminalNode
	WILDCARD() antlr.TerminalNode

	// IsExistsQueryContext differentiates from other interfaces.
	IsExistsQueryContext()
}

type ExistsQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExistsQueryContext() *ExistsQueryContext {
	var p = new(ExistsQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_existsQuery
	return p
}

func InitEmptyExistsQueryContext(p *ExistsQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_existsQuery
}

func (*ExistsQueryContext) IsExistsQueryContext() {}

func NewExistsQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExistsQueryContext {
	var p = new(ExistsQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_existsQuery

	return p
}

func (s *ExistsQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *ExistsQueryContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *ExistsQueryContext) COLON() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserCOLON, 0)
}

func (s *ExistsQueryContext) WILDCARD() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserWILDCARD, 0)
}

func (s *ExistsQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExistsQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExistsQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterExistsQuery(s)
	}
}

func (s *ExistsQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitExistsQuery(s)
	}
}

func (p *KqlBaseParser) ExistsQuery() (localctx IExistsQueryContext) {
	localctx = NewExistsQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, KqlBaseParserRULE_existsQuery)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.FieldName()
	}
	{
		p.SetState(122)
		p.Match(KqlBaseParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(KqlBaseParserWILDCARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldQueryContext is an interface to support dynamic dispatch.
type IFieldQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldName() IFieldNameContext
	COLON() antlr.TerminalNode
	FieldQueryValue() IFieldQueryValueContext

	// IsFieldQueryContext differentiates from other interfaces.
	IsFieldQueryContext()
}

type FieldQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldQueryContext() *FieldQueryContext {
	var p = new(FieldQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_fieldQuery
	return p
}

func InitEmptyFieldQueryContext(p *FieldQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_fieldQuery
}

func (*FieldQueryContext) IsFieldQueryContext() {}

func NewFieldQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldQueryContext {
	var p = new(FieldQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_fieldQuery

	return p
}

func (s *FieldQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldQueryContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldQueryContext) COLON() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserCOLON, 0)
}

func (s *FieldQueryContext) FieldQueryValue() IFieldQueryValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldQueryValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldQueryValueContext)
}

func (s *FieldQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterFieldQuery(s)
	}
}

func (s *FieldQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitFieldQuery(s)
	}
}

func (p *KqlBaseParser) FieldQuery() (localctx IFieldQueryContext) {
	localctx = NewFieldQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, KqlBaseParserRULE_fieldQuery)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.FieldName()
	}
	{
		p.SetState(126)
		p.Match(KqlBaseParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.FieldQueryValue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldLessQueryContext is an interface to support dynamic dispatch.
type IFieldLessQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldQueryValue() IFieldQueryValueContext

	// IsFieldLessQueryContext differentiates from other interfaces.
	IsFieldLessQueryContext()
}

type FieldLessQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldLessQueryContext() *FieldLessQueryContext {
	var p = new(FieldLessQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_fieldLessQuery
	return p
}

func InitEmptyFieldLessQueryContext(p *FieldLessQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_fieldLessQuery
}

func (*FieldLessQueryContext) IsFieldLessQueryContext() {}

func NewFieldLessQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldLessQueryContext {
	var p = new(FieldLessQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_fieldLessQuery

	return p
}

func (s *FieldLessQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldLessQueryContext) FieldQueryValue() IFieldQueryValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldQueryValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldQueryValueContext)
}

func (s *FieldLessQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldLessQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldLessQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterFieldLessQuery(s)
	}
}

func (s *FieldLessQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitFieldLessQuery(s)
	}
}

func (p *KqlBaseParser) FieldLessQuery() (localctx IFieldLessQueryContext) {
	localctx = NewFieldLessQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, KqlBaseParserRULE_fieldLessQuery)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.FieldQueryValue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldQueryValueContext is an interface to support dynamic dispatch.
type IFieldQueryValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the operator token.
	GetOperator() antlr.Token

	// SetOperator sets the operator token.
	SetOperator(antlr.Token)

	// Getter signatures
	FieldQueryValueLiteral() IFieldQueryValueLiteralContext
	FieldQueryValue() IFieldQueryValueContext
	NOT() antlr.TerminalNode
	LEFT_PARENTHESIS() antlr.TerminalNode
	BooleanFieldQueryValue() IBooleanFieldQueryValueContext
	RIGHT_PARENTHESIS() antlr.TerminalNode

	// IsFieldQueryValueContext differentiates from other interfaces.
	IsFieldQueryValueContext()
}

type FieldQueryValueContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	operator antlr.Token
}

func NewEmptyFieldQueryValueContext() *FieldQueryValueContext {
	var p = new(FieldQueryValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_fieldQueryValue
	return p
}

func InitEmptyFieldQueryValueContext(p *FieldQueryValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_fieldQueryValue
}

func (*FieldQueryValueContext) IsFieldQueryValueContext() {}

func NewFieldQueryValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldQueryValueContext {
	var p = new(FieldQueryValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_fieldQueryValue

	return p
}

func (s *FieldQueryValueContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldQueryValueContext) GetOperator() antlr.Token { return s.operator }

func (s *FieldQueryValueContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *FieldQueryValueContext) FieldQueryValueLiteral() IFieldQueryValueLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldQueryValueLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldQueryValueLiteralContext)
}

func (s *FieldQueryValueContext) FieldQueryValue() IFieldQueryValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldQueryValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldQueryValueContext)
}

func (s *FieldQueryValueContext) NOT() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserNOT, 0)
}

func (s *FieldQueryValueContext) LEFT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserLEFT_PARENTHESIS, 0)
}

func (s *FieldQueryValueContext) BooleanFieldQueryValue() IBooleanFieldQueryValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanFieldQueryValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanFieldQueryValueContext)
}

func (s *FieldQueryValueContext) RIGHT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserRIGHT_PARENTHESIS, 0)
}

func (s *FieldQueryValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldQueryValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldQueryValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterFieldQueryValue(s)
	}
}

func (s *FieldQueryValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitFieldQueryValue(s)
	}
}

func (p *KqlBaseParser) FieldQueryValue() (localctx IFieldQueryValueContext) {
	localctx = NewFieldQueryValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, KqlBaseParserRULE_fieldQueryValue)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.FieldQueryValueLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)

			var _m = p.Match(KqlBaseParserNOT)

			localctx.(*FieldQueryValueContext).operator = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.FieldQueryValue()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)
			p.Match(KqlBaseParserLEFT_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.booleanFieldQueryValue(0)
		}
		{
			p.SetState(136)
			p.Match(KqlBaseParserRIGHT_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldQueryValueLiteralContext is an interface to support dynamic dispatch.
type IFieldQueryValueLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldQueryValueUnquotedLiteral() IFieldQueryValueUnquotedLiteralContext
	QUOTED_STRING() antlr.TerminalNode

	// IsFieldQueryValueLiteralContext differentiates from other interfaces.
	IsFieldQueryValueLiteralContext()
}

type FieldQueryValueLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldQueryValueLiteralContext() *FieldQueryValueLiteralContext {
	var p = new(FieldQueryValueLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_fieldQueryValueLiteral
	return p
}

func InitEmptyFieldQueryValueLiteralContext(p *FieldQueryValueLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_fieldQueryValueLiteral
}

func (*FieldQueryValueLiteralContext) IsFieldQueryValueLiteralContext() {}

func NewFieldQueryValueLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldQueryValueLiteralContext {
	var p = new(FieldQueryValueLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_fieldQueryValueLiteral

	return p
}

func (s *FieldQueryValueLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldQueryValueLiteralContext) FieldQueryValueUnquotedLiteral() IFieldQueryValueUnquotedLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldQueryValueUnquotedLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldQueryValueUnquotedLiteralContext)
}

func (s *FieldQueryValueLiteralContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserQUOTED_STRING, 0)
}

func (s *FieldQueryValueLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldQueryValueLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldQueryValueLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterFieldQueryValueLiteral(s)
	}
}

func (s *FieldQueryValueLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitFieldQueryValueLiteral(s)
	}
}

func (p *KqlBaseParser) FieldQueryValueLiteral() (localctx IFieldQueryValueLiteralContext) {
	localctx = NewFieldQueryValueLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, KqlBaseParserRULE_fieldQueryValueLiteral)
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case KqlBaseParserAND, KqlBaseParserOR, KqlBaseParserNOT, KqlBaseParserUNQUOTED_LITERAL, KqlBaseParserWILDCARD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.FieldQueryValueUnquotedLiteral()
		}

	case KqlBaseParserQUOTED_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Match(KqlBaseParserQUOTED_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldQueryValueUnquotedLiteralContext is an interface to support dynamic dispatch.
type IFieldQueryValueUnquotedLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUNQUOTED_LITERAL() []antlr.TerminalNode
	UNQUOTED_LITERAL(i int) antlr.TerminalNode
	AllWILDCARD() []antlr.TerminalNode
	WILDCARD(i int) antlr.TerminalNode
	AllNOT() []antlr.TerminalNode
	NOT(i int) antlr.TerminalNode
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsFieldQueryValueUnquotedLiteralContext differentiates from other interfaces.
	IsFieldQueryValueUnquotedLiteralContext()
}

type FieldQueryValueUnquotedLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldQueryValueUnquotedLiteralContext() *FieldQueryValueUnquotedLiteralContext {
	var p = new(FieldQueryValueUnquotedLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_fieldQueryValueUnquotedLiteral
	return p
}

func InitEmptyFieldQueryValueUnquotedLiteralContext(p *FieldQueryValueUnquotedLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_fieldQueryValueUnquotedLiteral
}

func (*FieldQueryValueUnquotedLiteralContext) IsFieldQueryValueUnquotedLiteralContext() {}

func NewFieldQueryValueUnquotedLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldQueryValueUnquotedLiteralContext {
	var p = new(FieldQueryValueUnquotedLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_fieldQueryValueUnquotedLiteral

	return p
}

func (s *FieldQueryValueUnquotedLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldQueryValueUnquotedLiteralContext) AllUNQUOTED_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(KqlBaseParserUNQUOTED_LITERAL)
}

func (s *FieldQueryValueUnquotedLiteralContext) UNQUOTED_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(KqlBaseParserUNQUOTED_LITERAL, i)
}

func (s *FieldQueryValueUnquotedLiteralContext) AllWILDCARD() []antlr.TerminalNode {
	return s.GetTokens(KqlBaseParserWILDCARD)
}

func (s *FieldQueryValueUnquotedLiteralContext) WILDCARD(i int) antlr.TerminalNode {
	return s.GetToken(KqlBaseParserWILDCARD, i)
}

func (s *FieldQueryValueUnquotedLiteralContext) AllNOT() []antlr.TerminalNode {
	return s.GetTokens(KqlBaseParserNOT)
}

func (s *FieldQueryValueUnquotedLiteralContext) NOT(i int) antlr.TerminalNode {
	return s.GetToken(KqlBaseParserNOT, i)
}

func (s *FieldQueryValueUnquotedLiteralContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(KqlBaseParserAND)
}

func (s *FieldQueryValueUnquotedLiteralContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(KqlBaseParserAND, i)
}

func (s *FieldQueryValueUnquotedLiteralContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(KqlBaseParserOR)
}

func (s *FieldQueryValueUnquotedLiteralContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(KqlBaseParserOR, i)
}

func (s *FieldQueryValueUnquotedLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldQueryValueUnquotedLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldQueryValueUnquotedLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterFieldQueryValueUnquotedLiteral(s)
	}
}

func (s *FieldQueryValueUnquotedLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitFieldQueryValueUnquotedLiteral(s)
	}
}

func (p *KqlBaseParser) FieldQueryValueUnquotedLiteral() (localctx IFieldQueryValueUnquotedLiteralContext) {
	localctx = NewFieldQueryValueUnquotedLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, KqlBaseParserRULE_fieldQueryValueUnquotedLiteral)
	var _la int

	var _alt int

	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case KqlBaseParserUNQUOTED_LITERAL, KqlBaseParserWILDCARD:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(144)
					_la = p.GetTokenStream().LA(1)

					if !(_la == KqlBaseParserUNQUOTED_LITERAL || _la == KqlBaseParserWILDCARD) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(147)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(149)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&81936) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			}
			p.SetState(154)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(155)
				_la = p.GetTokenStream().LA(1)

				if !(_la == KqlBaseParserAND || _la == KqlBaseParserOR) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case KqlBaseParserAND, KqlBaseParserOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			_la = p.GetTokenStream().LA(1)

			if !(_la == KqlBaseParserAND || _la == KqlBaseParserOR) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			p.SetState(160)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(159)
						_la = p.GetTokenStream().LA(1)

						if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&81936) != 0) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}

				p.SetState(162)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}
			p.SetState(165)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(164)
					_la = p.GetTokenStream().LA(1)

					if !(_la == KqlBaseParserAND || _la == KqlBaseParserOR) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			} else if p.HasError() { // JIM
				goto errorExit
			}

		} else if p.HasError() { // JIM
			goto errorExit
		} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(167)
				p.Match(KqlBaseParserOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 3 {
			{
				p.SetState(168)
				p.Match(KqlBaseParserAND)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case KqlBaseParserNOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(171)
			p.Match(KqlBaseParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanFieldQueryValueContext is an interface to support dynamic dispatch.
type IBooleanFieldQueryValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the operator token.
	GetOperator() antlr.Token

	// SetOperator sets the operator token.
	SetOperator(antlr.Token)

	// Getter signatures
	LEFT_PARENTHESIS() antlr.TerminalNode
	BooleanFieldQueryValue() IBooleanFieldQueryValueContext
	RIGHT_PARENTHESIS() antlr.TerminalNode
	FieldQueryValue() IFieldQueryValueContext
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsBooleanFieldQueryValueContext differentiates from other interfaces.
	IsBooleanFieldQueryValueContext()
}

type BooleanFieldQueryValueContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	operator antlr.Token
}

func NewEmptyBooleanFieldQueryValueContext() *BooleanFieldQueryValueContext {
	var p = new(BooleanFieldQueryValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_booleanFieldQueryValue
	return p
}

func InitEmptyBooleanFieldQueryValueContext(p *BooleanFieldQueryValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_booleanFieldQueryValue
}

func (*BooleanFieldQueryValueContext) IsBooleanFieldQueryValueContext() {}

func NewBooleanFieldQueryValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanFieldQueryValueContext {
	var p = new(BooleanFieldQueryValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_booleanFieldQueryValue

	return p
}

func (s *BooleanFieldQueryValueContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanFieldQueryValueContext) GetOperator() antlr.Token { return s.operator }

func (s *BooleanFieldQueryValueContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *BooleanFieldQueryValueContext) LEFT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserLEFT_PARENTHESIS, 0)
}

func (s *BooleanFieldQueryValueContext) BooleanFieldQueryValue() IBooleanFieldQueryValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanFieldQueryValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanFieldQueryValueContext)
}

func (s *BooleanFieldQueryValueContext) RIGHT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserRIGHT_PARENTHESIS, 0)
}

func (s *BooleanFieldQueryValueContext) FieldQueryValue() IFieldQueryValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldQueryValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldQueryValueContext)
}

func (s *BooleanFieldQueryValueContext) AND() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserAND, 0)
}

func (s *BooleanFieldQueryValueContext) OR() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserOR, 0)
}

func (s *BooleanFieldQueryValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanFieldQueryValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanFieldQueryValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterBooleanFieldQueryValue(s)
	}
}

func (s *BooleanFieldQueryValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitBooleanFieldQueryValue(s)
	}
}

func (p *KqlBaseParser) BooleanFieldQueryValue() (localctx IBooleanFieldQueryValueContext) {
	return p.booleanFieldQueryValue(0)
}

func (p *KqlBaseParser) booleanFieldQueryValue(_p int) (localctx IBooleanFieldQueryValueContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewBooleanFieldQueryValueContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBooleanFieldQueryValueContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, KqlBaseParserRULE_booleanFieldQueryValue, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(175)
			p.Match(KqlBaseParserLEFT_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.booleanFieldQueryValue(0)
		}
		{
			p.SetState(177)
			p.Match(KqlBaseParserRIGHT_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(179)
			p.FieldQueryValue()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBooleanFieldQueryValueContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, KqlBaseParserRULE_booleanFieldQueryValue)
			p.SetState(182)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(183)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*BooleanFieldQueryValueContext).operator = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == KqlBaseParserAND || _la == KqlBaseParserOR) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*BooleanFieldQueryValueContext).operator = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(184)
				p.FieldQueryValue()
			}

		}
		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// Getter signatures
	UNQUOTED_LITERAL() antlr.TerminalNode
	QUOTED_STRING() antlr.TerminalNode
	WILDCARD() antlr.TerminalNode

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_fieldName
	return p
}

func InitEmptyFieldNameContext(p *FieldNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KqlBaseParserRULE_fieldName
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KqlBaseParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) GetValue() antlr.Token { return s.value }

func (s *FieldNameContext) SetValue(v antlr.Token) { s.value = v }

func (s *FieldNameContext) UNQUOTED_LITERAL() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserUNQUOTED_LITERAL, 0)
}

func (s *FieldNameContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserQUOTED_STRING, 0)
}

func (s *FieldNameContext) WILDCARD() antlr.TerminalNode {
	return s.GetToken(KqlBaseParserWILDCARD, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KqlBaseListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *KqlBaseParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, KqlBaseParserRULE_fieldName)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case KqlBaseParserUNQUOTED_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)

			var _m = p.Match(KqlBaseParserUNQUOTED_LITERAL)

			localctx.(*FieldNameContext).value = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case KqlBaseParserQUOTED_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)

			var _m = p.Match(KqlBaseParserQUOTED_STRING)

			localctx.(*FieldNameContext).value = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case KqlBaseParserWILDCARD:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(192)

			var _m = p.Match(KqlBaseParserWILDCARD)

			localctx.(*FieldNameContext).value = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *KqlBaseParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *QueryContext = nil
		if localctx != nil {
			t = localctx.(*QueryContext)
		}
		return p.Query_Sempred(t, predIndex)

	case 5:
		var t *NestedSubQueryContext = nil
		if localctx != nil {
			t = localctx.(*NestedSubQueryContext)
		}
		return p.NestedSubQuery_Sempred(t, predIndex)

	case 18:
		var t *BooleanFieldQueryValueContext = nil
		if localctx != nil {
			t = localctx.(*BooleanFieldQueryValueContext)
		}
		return p.BooleanFieldQueryValue_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *KqlBaseParser) Query_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *KqlBaseParser) NestedSubQuery_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *KqlBaseParser) BooleanFieldQueryValue_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
