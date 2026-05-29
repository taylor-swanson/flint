// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package esql

import "github.com/antlr4-go/antlr/v4"

type LexerConfig struct {
	*antlr.BaseLexer
	promqlDepth int
}

func (l *LexerConfig) incPromqlDepth() {
	l.promqlDepth++
}

func (l *LexerConfig) decPromqlDepth() {
	if l.promqlDepth > 0 {
		l.promqlDepth--
	}
}

func (l *LexerConfig) isPromqlQuery() bool {
	return l.promqlDepth > 0
}

func (l *LexerConfig) resetPromqlDepth() {
	l.promqlDepth = 0
}

func (l *LexerConfig) isDevVersion() bool {
	return false
}

type ParserConfig struct {
	*antlr.BaseParser
}

func (p *ParserConfig) isDevVersion() bool {
	return false
}
