// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// Code generated from EsqlBaseParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package esql // EsqlBaseParser
import "github.com/antlr4-go/antlr/v4"

// EsqlBaseParserListener is a complete listener for a parse tree produced by EsqlBaseParser.
type EsqlBaseParserListener interface {
	antlr.ParseTreeListener

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterSingleStatement is called when entering the singleStatement production.
	EnterSingleStatement(c *SingleStatementContext)

	// EnterCompositeQuery is called when entering the compositeQuery production.
	EnterCompositeQuery(c *CompositeQueryContext)

	// EnterSingleCommandQuery is called when entering the singleCommandQuery production.
	EnterSingleCommandQuery(c *SingleCommandQueryContext)

	// EnterSourceCommand is called when entering the sourceCommand production.
	EnterSourceCommand(c *SourceCommandContext)

	// EnterProcessingCommand is called when entering the processingCommand production.
	EnterProcessingCommand(c *ProcessingCommandContext)

	// EnterWhereCommand is called when entering the whereCommand production.
	EnterWhereCommand(c *WhereCommandContext)

	// EnterToDataType is called when entering the toDataType production.
	EnterToDataType(c *ToDataTypeContext)

	// EnterRowCommand is called when entering the rowCommand production.
	EnterRowCommand(c *RowCommandContext)

	// EnterFields is called when entering the fields production.
	EnterFields(c *FieldsContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFromCommand is called when entering the fromCommand production.
	EnterFromCommand(c *FromCommandContext)

	// EnterTimeSeriesCommand is called when entering the timeSeriesCommand production.
	EnterTimeSeriesCommand(c *TimeSeriesCommandContext)

	// EnterExternalCommand is called when entering the externalCommand production.
	EnterExternalCommand(c *ExternalCommandContext)

	// EnterIndexPatternAndMetadataFields is called when entering the indexPatternAndMetadataFields production.
	EnterIndexPatternAndMetadataFields(c *IndexPatternAndMetadataFieldsContext)

	// EnterIndexPatternOrSubquery is called when entering the indexPatternOrSubquery production.
	EnterIndexPatternOrSubquery(c *IndexPatternOrSubqueryContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

	// EnterIndexPattern is called when entering the indexPattern production.
	EnterIndexPattern(c *IndexPatternContext)

	// EnterClusterString is called when entering the clusterString production.
	EnterClusterString(c *ClusterStringContext)

	// EnterSelectorString is called when entering the selectorString production.
	EnterSelectorString(c *SelectorStringContext)

	// EnterUnquotedIndexString is called when entering the unquotedIndexString production.
	EnterUnquotedIndexString(c *UnquotedIndexStringContext)

	// EnterIndexString is called when entering the indexString production.
	EnterIndexString(c *IndexStringContext)

	// EnterMetadata is called when entering the metadata production.
	EnterMetadata(c *MetadataContext)

	// EnterEvalCommand is called when entering the evalCommand production.
	EnterEvalCommand(c *EvalCommandContext)

	// EnterStatsCommand is called when entering the statsCommand production.
	EnterStatsCommand(c *StatsCommandContext)

	// EnterAggFields is called when entering the aggFields production.
	EnterAggFields(c *AggFieldsContext)

	// EnterAggField is called when entering the aggField production.
	EnterAggField(c *AggFieldContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterQualifiedNamePattern is called when entering the qualifiedNamePattern production.
	EnterQualifiedNamePattern(c *QualifiedNamePatternContext)

	// EnterFieldNamePattern is called when entering the fieldNamePattern production.
	EnterFieldNamePattern(c *FieldNamePatternContext)

	// EnterQualifiedNamePatterns is called when entering the qualifiedNamePatterns production.
	EnterQualifiedNamePatterns(c *QualifiedNamePatternsContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterIdentifierPattern is called when entering the identifierPattern production.
	EnterIdentifierPattern(c *IdentifierPatternContext)

	// EnterInputParam is called when entering the inputParam production.
	EnterInputParam(c *InputParamContext)

	// EnterInputNamedOrPositionalParam is called when entering the inputNamedOrPositionalParam production.
	EnterInputNamedOrPositionalParam(c *InputNamedOrPositionalParamContext)

	// EnterInputDoubleParams is called when entering the inputDoubleParams production.
	EnterInputDoubleParams(c *InputDoubleParamsContext)

	// EnterInputNamedOrPositionalDoubleParams is called when entering the inputNamedOrPositionalDoubleParams production.
	EnterInputNamedOrPositionalDoubleParams(c *InputNamedOrPositionalDoubleParamsContext)

	// EnterIdentifierOrParameter is called when entering the identifierOrParameter production.
	EnterIdentifierOrParameter(c *IdentifierOrParameterContext)

	// EnterStringOrParameter is called when entering the stringOrParameter production.
	EnterStringOrParameter(c *StringOrParameterContext)

	// EnterLimitCommand is called when entering the limitCommand production.
	EnterLimitCommand(c *LimitCommandContext)

	// EnterLimitByGroupKey is called when entering the limitByGroupKey production.
	EnterLimitByGroupKey(c *LimitByGroupKeyContext)

	// EnterSortCommand is called when entering the sortCommand production.
	EnterSortCommand(c *SortCommandContext)

	// EnterOrderExpression is called when entering the orderExpression production.
	EnterOrderExpression(c *OrderExpressionContext)

	// EnterKeepCommand is called when entering the keepCommand production.
	EnterKeepCommand(c *KeepCommandContext)

	// EnterDropCommand is called when entering the dropCommand production.
	EnterDropCommand(c *DropCommandContext)

	// EnterRenameCommand is called when entering the renameCommand production.
	EnterRenameCommand(c *RenameCommandContext)

	// EnterRenameClause is called when entering the renameClause production.
	EnterRenameClause(c *RenameClauseContext)

	// EnterDissectCommand is called when entering the dissectCommand production.
	EnterDissectCommand(c *DissectCommandContext)

	// EnterDissectCommandOptions is called when entering the dissectCommandOptions production.
	EnterDissectCommandOptions(c *DissectCommandOptionsContext)

	// EnterDissectCommandOption is called when entering the dissectCommandOption production.
	EnterDissectCommandOption(c *DissectCommandOptionContext)

	// EnterCommandNamedParameters is called when entering the commandNamedParameters production.
	EnterCommandNamedParameters(c *CommandNamedParametersContext)

	// EnterGrokCommand is called when entering the grokCommand production.
	EnterGrokCommand(c *GrokCommandContext)

	// EnterMvExpandCommand is called when entering the mvExpandCommand production.
	EnterMvExpandCommand(c *MvExpandCommandContext)

	// EnterExplainCommand is called when entering the explainCommand production.
	EnterExplainCommand(c *ExplainCommandContext)

	// EnterSubqueryExpression is called when entering the subqueryExpression production.
	EnterSubqueryExpression(c *SubqueryExpressionContext)

	// EnterShowInfo is called when entering the showInfo production.
	EnterShowInfo(c *ShowInfoContext)

	// EnterEnrichCommand is called when entering the enrichCommand production.
	EnterEnrichCommand(c *EnrichCommandContext)

	// EnterEnrichPolicyName is called when entering the enrichPolicyName production.
	EnterEnrichPolicyName(c *EnrichPolicyNameContext)

	// EnterEnrichWithClause is called when entering the enrichWithClause production.
	EnterEnrichWithClause(c *EnrichWithClauseContext)

	// EnterSampleCommand is called when entering the sampleCommand production.
	EnterSampleCommand(c *SampleCommandContext)

	// EnterChangePointCommand is called when entering the changePointCommand production.
	EnterChangePointCommand(c *ChangePointCommandContext)

	// EnterForkCommand is called when entering the forkCommand production.
	EnterForkCommand(c *ForkCommandContext)

	// EnterForkSubQueries is called when entering the forkSubQueries production.
	EnterForkSubQueries(c *ForkSubQueriesContext)

	// EnterForkSubQuery is called when entering the forkSubQuery production.
	EnterForkSubQuery(c *ForkSubQueryContext)

	// EnterSingleForkSubQueryCommand is called when entering the singleForkSubQueryCommand production.
	EnterSingleForkSubQueryCommand(c *SingleForkSubQueryCommandContext)

	// EnterCompositeForkSubQuery is called when entering the compositeForkSubQuery production.
	EnterCompositeForkSubQuery(c *CompositeForkSubQueryContext)

	// EnterForkSubQueryProcessingCommand is called when entering the forkSubQueryProcessingCommand production.
	EnterForkSubQueryProcessingCommand(c *ForkSubQueryProcessingCommandContext)

	// EnterRerankCommand is called when entering the rerankCommand production.
	EnterRerankCommand(c *RerankCommandContext)

	// EnterCompletionCommand is called when entering the completionCommand production.
	EnterCompletionCommand(c *CompletionCommandContext)

	// EnterInlineStatsCommand is called when entering the inlineStatsCommand production.
	EnterInlineStatsCommand(c *InlineStatsCommandContext)

	// EnterFuseCommand is called when entering the fuseCommand production.
	EnterFuseCommand(c *FuseCommandContext)

	// EnterFuseConfiguration is called when entering the fuseConfiguration production.
	EnterFuseConfiguration(c *FuseConfigurationContext)

	// EnterFuseKeyByFields is called when entering the fuseKeyByFields production.
	EnterFuseKeyByFields(c *FuseKeyByFieldsContext)

	// EnterMetricsInfoCommand is called when entering the metricsInfoCommand production.
	EnterMetricsInfoCommand(c *MetricsInfoCommandContext)

	// EnterTsInfoCommand is called when entering the tsInfoCommand production.
	EnterTsInfoCommand(c *TsInfoCommandContext)

	// EnterLookupCommand is called when entering the lookupCommand production.
	EnterLookupCommand(c *LookupCommandContext)

	// EnterInsistCommand is called when entering the insistCommand production.
	EnterInsistCommand(c *InsistCommandContext)

	// EnterUriPartsCommand is called when entering the uriPartsCommand production.
	EnterUriPartsCommand(c *UriPartsCommandContext)

	// EnterRegisteredDomainCommand is called when entering the registeredDomainCommand production.
	EnterRegisteredDomainCommand(c *RegisteredDomainCommandContext)

	// EnterSetCommand is called when entering the setCommand production.
	EnterSetCommand(c *SetCommandContext)

	// EnterSetField is called when entering the setField production.
	EnterSetField(c *SetFieldContext)

	// EnterMmrCommand is called when entering the mmrCommand production.
	EnterMmrCommand(c *MmrCommandContext)

	// EnterMmrQueryVectorParameter is called when entering the mmrQueryVectorParameter production.
	EnterMmrQueryVectorParameter(c *MmrQueryVectorParameterContext)

	// EnterMmrQueryVectorExpression is called when entering the mmrQueryVectorExpression production.
	EnterMmrQueryVectorExpression(c *MmrQueryVectorExpressionContext)

	// EnterMatchExpression is called when entering the matchExpression production.
	EnterMatchExpression(c *MatchExpressionContext)

	// EnterLogicalNot is called when entering the logicalNot production.
	EnterLogicalNot(c *LogicalNotContext)

	// EnterBooleanDefault is called when entering the booleanDefault production.
	EnterBooleanDefault(c *BooleanDefaultContext)

	// EnterIsNull is called when entering the isNull production.
	EnterIsNull(c *IsNullContext)

	// EnterRegexExpression is called when entering the regexExpression production.
	EnterRegexExpression(c *RegexExpressionContext)

	// EnterLogicalIn is called when entering the logicalIn production.
	EnterLogicalIn(c *LogicalInContext)

	// EnterLogicalBinary is called when entering the logicalBinary production.
	EnterLogicalBinary(c *LogicalBinaryContext)

	// EnterLikeExpression is called when entering the likeExpression production.
	EnterLikeExpression(c *LikeExpressionContext)

	// EnterRlikeExpression is called when entering the rlikeExpression production.
	EnterRlikeExpression(c *RlikeExpressionContext)

	// EnterLikeListExpression is called when entering the likeListExpression production.
	EnterLikeListExpression(c *LikeListExpressionContext)

	// EnterRlikeListExpression is called when entering the rlikeListExpression production.
	EnterRlikeListExpression(c *RlikeListExpressionContext)

	// EnterMatchBooleanExpression is called when entering the matchBooleanExpression production.
	EnterMatchBooleanExpression(c *MatchBooleanExpressionContext)

	// EnterValueExpressionDefault is called when entering the valueExpressionDefault production.
	EnterValueExpressionDefault(c *ValueExpressionDefaultContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterOperatorExpressionDefault is called when entering the operatorExpressionDefault production.
	EnterOperatorExpressionDefault(c *OperatorExpressionDefaultContext)

	// EnterArithmeticBinary is called when entering the arithmeticBinary production.
	EnterArithmeticBinary(c *ArithmeticBinaryContext)

	// EnterArithmeticUnary is called when entering the arithmeticUnary production.
	EnterArithmeticUnary(c *ArithmeticUnaryContext)

	// EnterDereference is called when entering the dereference production.
	EnterDereference(c *DereferenceContext)

	// EnterInlineCast is called when entering the inlineCast production.
	EnterInlineCast(c *InlineCastContext)

	// EnterConstantDefault is called when entering the constantDefault production.
	EnterConstantDefault(c *ConstantDefaultContext)

	// EnterParenthesizedExpression is called when entering the parenthesizedExpression production.
	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFunctionExpression is called when entering the functionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterMapExpression is called when entering the mapExpression production.
	EnterMapExpression(c *MapExpressionContext)

	// EnterEntryExpression is called when entering the entryExpression production.
	EnterEntryExpression(c *EntryExpressionContext)

	// EnterMapValue is called when entering the mapValue production.
	EnterMapValue(c *MapValueContext)

	// EnterNullLiteral is called when entering the nullLiteral production.
	EnterNullLiteral(c *NullLiteralContext)

	// EnterQualifiedIntegerLiteral is called when entering the qualifiedIntegerLiteral production.
	EnterQualifiedIntegerLiteral(c *QualifiedIntegerLiteralContext)

	// EnterDecimalLiteral is called when entering the decimalLiteral production.
	EnterDecimalLiteral(c *DecimalLiteralContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterInputParameter is called when entering the inputParameter production.
	EnterInputParameter(c *InputParameterContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterNumericArrayLiteral is called when entering the numericArrayLiteral production.
	EnterNumericArrayLiteral(c *NumericArrayLiteralContext)

	// EnterBooleanArrayLiteral is called when entering the booleanArrayLiteral production.
	EnterBooleanArrayLiteral(c *BooleanArrayLiteralContext)

	// EnterStringArrayLiteral is called when entering the stringArrayLiteral production.
	EnterStringArrayLiteral(c *StringArrayLiteralContext)

	// EnterBooleanValue is called when entering the booleanValue production.
	EnterBooleanValue(c *BooleanValueContext)

	// EnterNumericValue is called when entering the numericValue production.
	EnterNumericValue(c *NumericValueContext)

	// EnterDecimalValue is called when entering the decimalValue production.
	EnterDecimalValue(c *DecimalValueContext)

	// EnterIntegerValue is called when entering the integerValue production.
	EnterIntegerValue(c *IntegerValueContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterJoinCommand is called when entering the joinCommand production.
	EnterJoinCommand(c *JoinCommandContext)

	// EnterJoinTarget is called when entering the joinTarget production.
	EnterJoinTarget(c *JoinTargetContext)

	// EnterJoinCondition is called when entering the joinCondition production.
	EnterJoinCondition(c *JoinConditionContext)

	// EnterPromqlCommand is called when entering the promqlCommand production.
	EnterPromqlCommand(c *PromqlCommandContext)

	// EnterValueName is called when entering the valueName production.
	EnterValueName(c *ValueNameContext)

	// EnterPromqlParam is called when entering the promqlParam production.
	EnterPromqlParam(c *PromqlParamContext)

	// EnterPromqlParamName is called when entering the promqlParamName production.
	EnterPromqlParamName(c *PromqlParamNameContext)

	// EnterPromqlParamValue is called when entering the promqlParamValue production.
	EnterPromqlParamValue(c *PromqlParamValueContext)

	// EnterPromqlQueryContent is called when entering the promqlQueryContent production.
	EnterPromqlQueryContent(c *PromqlQueryContentContext)

	// EnterPromqlQueryPart is called when entering the promqlQueryPart production.
	EnterPromqlQueryPart(c *PromqlQueryPartContext)

	// EnterPromqlIndexPattern is called when entering the promqlIndexPattern production.
	EnterPromqlIndexPattern(c *PromqlIndexPatternContext)

	// EnterPromqlClusterString is called when entering the promqlClusterString production.
	EnterPromqlClusterString(c *PromqlClusterStringContext)

	// EnterPromqlSelectorString is called when entering the promqlSelectorString production.
	EnterPromqlSelectorString(c *PromqlSelectorStringContext)

	// EnterPromqlUnquotedIndexString is called when entering the promqlUnquotedIndexString production.
	EnterPromqlUnquotedIndexString(c *PromqlUnquotedIndexStringContext)

	// EnterPromqlIndexString is called when entering the promqlIndexString production.
	EnterPromqlIndexString(c *PromqlIndexStringContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitSingleStatement is called when exiting the singleStatement production.
	ExitSingleStatement(c *SingleStatementContext)

	// ExitCompositeQuery is called when exiting the compositeQuery production.
	ExitCompositeQuery(c *CompositeQueryContext)

	// ExitSingleCommandQuery is called when exiting the singleCommandQuery production.
	ExitSingleCommandQuery(c *SingleCommandQueryContext)

	// ExitSourceCommand is called when exiting the sourceCommand production.
	ExitSourceCommand(c *SourceCommandContext)

	// ExitProcessingCommand is called when exiting the processingCommand production.
	ExitProcessingCommand(c *ProcessingCommandContext)

	// ExitWhereCommand is called when exiting the whereCommand production.
	ExitWhereCommand(c *WhereCommandContext)

	// ExitToDataType is called when exiting the toDataType production.
	ExitToDataType(c *ToDataTypeContext)

	// ExitRowCommand is called when exiting the rowCommand production.
	ExitRowCommand(c *RowCommandContext)

	// ExitFields is called when exiting the fields production.
	ExitFields(c *FieldsContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFromCommand is called when exiting the fromCommand production.
	ExitFromCommand(c *FromCommandContext)

	// ExitTimeSeriesCommand is called when exiting the timeSeriesCommand production.
	ExitTimeSeriesCommand(c *TimeSeriesCommandContext)

	// ExitExternalCommand is called when exiting the externalCommand production.
	ExitExternalCommand(c *ExternalCommandContext)

	// ExitIndexPatternAndMetadataFields is called when exiting the indexPatternAndMetadataFields production.
	ExitIndexPatternAndMetadataFields(c *IndexPatternAndMetadataFieldsContext)

	// ExitIndexPatternOrSubquery is called when exiting the indexPatternOrSubquery production.
	ExitIndexPatternOrSubquery(c *IndexPatternOrSubqueryContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

	// ExitIndexPattern is called when exiting the indexPattern production.
	ExitIndexPattern(c *IndexPatternContext)

	// ExitClusterString is called when exiting the clusterString production.
	ExitClusterString(c *ClusterStringContext)

	// ExitSelectorString is called when exiting the selectorString production.
	ExitSelectorString(c *SelectorStringContext)

	// ExitUnquotedIndexString is called when exiting the unquotedIndexString production.
	ExitUnquotedIndexString(c *UnquotedIndexStringContext)

	// ExitIndexString is called when exiting the indexString production.
	ExitIndexString(c *IndexStringContext)

	// ExitMetadata is called when exiting the metadata production.
	ExitMetadata(c *MetadataContext)

	// ExitEvalCommand is called when exiting the evalCommand production.
	ExitEvalCommand(c *EvalCommandContext)

	// ExitStatsCommand is called when exiting the statsCommand production.
	ExitStatsCommand(c *StatsCommandContext)

	// ExitAggFields is called when exiting the aggFields production.
	ExitAggFields(c *AggFieldsContext)

	// ExitAggField is called when exiting the aggField production.
	ExitAggField(c *AggFieldContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitQualifiedNamePattern is called when exiting the qualifiedNamePattern production.
	ExitQualifiedNamePattern(c *QualifiedNamePatternContext)

	// ExitFieldNamePattern is called when exiting the fieldNamePattern production.
	ExitFieldNamePattern(c *FieldNamePatternContext)

	// ExitQualifiedNamePatterns is called when exiting the qualifiedNamePatterns production.
	ExitQualifiedNamePatterns(c *QualifiedNamePatternsContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitIdentifierPattern is called when exiting the identifierPattern production.
	ExitIdentifierPattern(c *IdentifierPatternContext)

	// ExitInputParam is called when exiting the inputParam production.
	ExitInputParam(c *InputParamContext)

	// ExitInputNamedOrPositionalParam is called when exiting the inputNamedOrPositionalParam production.
	ExitInputNamedOrPositionalParam(c *InputNamedOrPositionalParamContext)

	// ExitInputDoubleParams is called when exiting the inputDoubleParams production.
	ExitInputDoubleParams(c *InputDoubleParamsContext)

	// ExitInputNamedOrPositionalDoubleParams is called when exiting the inputNamedOrPositionalDoubleParams production.
	ExitInputNamedOrPositionalDoubleParams(c *InputNamedOrPositionalDoubleParamsContext)

	// ExitIdentifierOrParameter is called when exiting the identifierOrParameter production.
	ExitIdentifierOrParameter(c *IdentifierOrParameterContext)

	// ExitStringOrParameter is called when exiting the stringOrParameter production.
	ExitStringOrParameter(c *StringOrParameterContext)

	// ExitLimitCommand is called when exiting the limitCommand production.
	ExitLimitCommand(c *LimitCommandContext)

	// ExitLimitByGroupKey is called when exiting the limitByGroupKey production.
	ExitLimitByGroupKey(c *LimitByGroupKeyContext)

	// ExitSortCommand is called when exiting the sortCommand production.
	ExitSortCommand(c *SortCommandContext)

	// ExitOrderExpression is called when exiting the orderExpression production.
	ExitOrderExpression(c *OrderExpressionContext)

	// ExitKeepCommand is called when exiting the keepCommand production.
	ExitKeepCommand(c *KeepCommandContext)

	// ExitDropCommand is called when exiting the dropCommand production.
	ExitDropCommand(c *DropCommandContext)

	// ExitRenameCommand is called when exiting the renameCommand production.
	ExitRenameCommand(c *RenameCommandContext)

	// ExitRenameClause is called when exiting the renameClause production.
	ExitRenameClause(c *RenameClauseContext)

	// ExitDissectCommand is called when exiting the dissectCommand production.
	ExitDissectCommand(c *DissectCommandContext)

	// ExitDissectCommandOptions is called when exiting the dissectCommandOptions production.
	ExitDissectCommandOptions(c *DissectCommandOptionsContext)

	// ExitDissectCommandOption is called when exiting the dissectCommandOption production.
	ExitDissectCommandOption(c *DissectCommandOptionContext)

	// ExitCommandNamedParameters is called when exiting the commandNamedParameters production.
	ExitCommandNamedParameters(c *CommandNamedParametersContext)

	// ExitGrokCommand is called when exiting the grokCommand production.
	ExitGrokCommand(c *GrokCommandContext)

	// ExitMvExpandCommand is called when exiting the mvExpandCommand production.
	ExitMvExpandCommand(c *MvExpandCommandContext)

	// ExitExplainCommand is called when exiting the explainCommand production.
	ExitExplainCommand(c *ExplainCommandContext)

	// ExitSubqueryExpression is called when exiting the subqueryExpression production.
	ExitSubqueryExpression(c *SubqueryExpressionContext)

	// ExitShowInfo is called when exiting the showInfo production.
	ExitShowInfo(c *ShowInfoContext)

	// ExitEnrichCommand is called when exiting the enrichCommand production.
	ExitEnrichCommand(c *EnrichCommandContext)

	// ExitEnrichPolicyName is called when exiting the enrichPolicyName production.
	ExitEnrichPolicyName(c *EnrichPolicyNameContext)

	// ExitEnrichWithClause is called when exiting the enrichWithClause production.
	ExitEnrichWithClause(c *EnrichWithClauseContext)

	// ExitSampleCommand is called when exiting the sampleCommand production.
	ExitSampleCommand(c *SampleCommandContext)

	// ExitChangePointCommand is called when exiting the changePointCommand production.
	ExitChangePointCommand(c *ChangePointCommandContext)

	// ExitForkCommand is called when exiting the forkCommand production.
	ExitForkCommand(c *ForkCommandContext)

	// ExitForkSubQueries is called when exiting the forkSubQueries production.
	ExitForkSubQueries(c *ForkSubQueriesContext)

	// ExitForkSubQuery is called when exiting the forkSubQuery production.
	ExitForkSubQuery(c *ForkSubQueryContext)

	// ExitSingleForkSubQueryCommand is called when exiting the singleForkSubQueryCommand production.
	ExitSingleForkSubQueryCommand(c *SingleForkSubQueryCommandContext)

	// ExitCompositeForkSubQuery is called when exiting the compositeForkSubQuery production.
	ExitCompositeForkSubQuery(c *CompositeForkSubQueryContext)

	// ExitForkSubQueryProcessingCommand is called when exiting the forkSubQueryProcessingCommand production.
	ExitForkSubQueryProcessingCommand(c *ForkSubQueryProcessingCommandContext)

	// ExitRerankCommand is called when exiting the rerankCommand production.
	ExitRerankCommand(c *RerankCommandContext)

	// ExitCompletionCommand is called when exiting the completionCommand production.
	ExitCompletionCommand(c *CompletionCommandContext)

	// ExitInlineStatsCommand is called when exiting the inlineStatsCommand production.
	ExitInlineStatsCommand(c *InlineStatsCommandContext)

	// ExitFuseCommand is called when exiting the fuseCommand production.
	ExitFuseCommand(c *FuseCommandContext)

	// ExitFuseConfiguration is called when exiting the fuseConfiguration production.
	ExitFuseConfiguration(c *FuseConfigurationContext)

	// ExitFuseKeyByFields is called when exiting the fuseKeyByFields production.
	ExitFuseKeyByFields(c *FuseKeyByFieldsContext)

	// ExitMetricsInfoCommand is called when exiting the metricsInfoCommand production.
	ExitMetricsInfoCommand(c *MetricsInfoCommandContext)

	// ExitTsInfoCommand is called when exiting the tsInfoCommand production.
	ExitTsInfoCommand(c *TsInfoCommandContext)

	// ExitLookupCommand is called when exiting the lookupCommand production.
	ExitLookupCommand(c *LookupCommandContext)

	// ExitInsistCommand is called when exiting the insistCommand production.
	ExitInsistCommand(c *InsistCommandContext)

	// ExitUriPartsCommand is called when exiting the uriPartsCommand production.
	ExitUriPartsCommand(c *UriPartsCommandContext)

	// ExitRegisteredDomainCommand is called when exiting the registeredDomainCommand production.
	ExitRegisteredDomainCommand(c *RegisteredDomainCommandContext)

	// ExitSetCommand is called when exiting the setCommand production.
	ExitSetCommand(c *SetCommandContext)

	// ExitSetField is called when exiting the setField production.
	ExitSetField(c *SetFieldContext)

	// ExitMmrCommand is called when exiting the mmrCommand production.
	ExitMmrCommand(c *MmrCommandContext)

	// ExitMmrQueryVectorParameter is called when exiting the mmrQueryVectorParameter production.
	ExitMmrQueryVectorParameter(c *MmrQueryVectorParameterContext)

	// ExitMmrQueryVectorExpression is called when exiting the mmrQueryVectorExpression production.
	ExitMmrQueryVectorExpression(c *MmrQueryVectorExpressionContext)

	// ExitMatchExpression is called when exiting the matchExpression production.
	ExitMatchExpression(c *MatchExpressionContext)

	// ExitLogicalNot is called when exiting the logicalNot production.
	ExitLogicalNot(c *LogicalNotContext)

	// ExitBooleanDefault is called when exiting the booleanDefault production.
	ExitBooleanDefault(c *BooleanDefaultContext)

	// ExitIsNull is called when exiting the isNull production.
	ExitIsNull(c *IsNullContext)

	// ExitRegexExpression is called when exiting the regexExpression production.
	ExitRegexExpression(c *RegexExpressionContext)

	// ExitLogicalIn is called when exiting the logicalIn production.
	ExitLogicalIn(c *LogicalInContext)

	// ExitLogicalBinary is called when exiting the logicalBinary production.
	ExitLogicalBinary(c *LogicalBinaryContext)

	// ExitLikeExpression is called when exiting the likeExpression production.
	ExitLikeExpression(c *LikeExpressionContext)

	// ExitRlikeExpression is called when exiting the rlikeExpression production.
	ExitRlikeExpression(c *RlikeExpressionContext)

	// ExitLikeListExpression is called when exiting the likeListExpression production.
	ExitLikeListExpression(c *LikeListExpressionContext)

	// ExitRlikeListExpression is called when exiting the rlikeListExpression production.
	ExitRlikeListExpression(c *RlikeListExpressionContext)

	// ExitMatchBooleanExpression is called when exiting the matchBooleanExpression production.
	ExitMatchBooleanExpression(c *MatchBooleanExpressionContext)

	// ExitValueExpressionDefault is called when exiting the valueExpressionDefault production.
	ExitValueExpressionDefault(c *ValueExpressionDefaultContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitOperatorExpressionDefault is called when exiting the operatorExpressionDefault production.
	ExitOperatorExpressionDefault(c *OperatorExpressionDefaultContext)

	// ExitArithmeticBinary is called when exiting the arithmeticBinary production.
	ExitArithmeticBinary(c *ArithmeticBinaryContext)

	// ExitArithmeticUnary is called when exiting the arithmeticUnary production.
	ExitArithmeticUnary(c *ArithmeticUnaryContext)

	// ExitDereference is called when exiting the dereference production.
	ExitDereference(c *DereferenceContext)

	// ExitInlineCast is called when exiting the inlineCast production.
	ExitInlineCast(c *InlineCastContext)

	// ExitConstantDefault is called when exiting the constantDefault production.
	ExitConstantDefault(c *ConstantDefaultContext)

	// ExitParenthesizedExpression is called when exiting the parenthesizedExpression production.
	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFunctionExpression is called when exiting the functionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitMapExpression is called when exiting the mapExpression production.
	ExitMapExpression(c *MapExpressionContext)

	// ExitEntryExpression is called when exiting the entryExpression production.
	ExitEntryExpression(c *EntryExpressionContext)

	// ExitMapValue is called when exiting the mapValue production.
	ExitMapValue(c *MapValueContext)

	// ExitNullLiteral is called when exiting the nullLiteral production.
	ExitNullLiteral(c *NullLiteralContext)

	// ExitQualifiedIntegerLiteral is called when exiting the qualifiedIntegerLiteral production.
	ExitQualifiedIntegerLiteral(c *QualifiedIntegerLiteralContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *DecimalLiteralContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitInputParameter is called when exiting the inputParameter production.
	ExitInputParameter(c *InputParameterContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitNumericArrayLiteral is called when exiting the numericArrayLiteral production.
	ExitNumericArrayLiteral(c *NumericArrayLiteralContext)

	// ExitBooleanArrayLiteral is called when exiting the booleanArrayLiteral production.
	ExitBooleanArrayLiteral(c *BooleanArrayLiteralContext)

	// ExitStringArrayLiteral is called when exiting the stringArrayLiteral production.
	ExitStringArrayLiteral(c *StringArrayLiteralContext)

	// ExitBooleanValue is called when exiting the booleanValue production.
	ExitBooleanValue(c *BooleanValueContext)

	// ExitNumericValue is called when exiting the numericValue production.
	ExitNumericValue(c *NumericValueContext)

	// ExitDecimalValue is called when exiting the decimalValue production.
	ExitDecimalValue(c *DecimalValueContext)

	// ExitIntegerValue is called when exiting the integerValue production.
	ExitIntegerValue(c *IntegerValueContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitJoinCommand is called when exiting the joinCommand production.
	ExitJoinCommand(c *JoinCommandContext)

	// ExitJoinTarget is called when exiting the joinTarget production.
	ExitJoinTarget(c *JoinTargetContext)

	// ExitJoinCondition is called when exiting the joinCondition production.
	ExitJoinCondition(c *JoinConditionContext)

	// ExitPromqlCommand is called when exiting the promqlCommand production.
	ExitPromqlCommand(c *PromqlCommandContext)

	// ExitValueName is called when exiting the valueName production.
	ExitValueName(c *ValueNameContext)

	// ExitPromqlParam is called when exiting the promqlParam production.
	ExitPromqlParam(c *PromqlParamContext)

	// ExitPromqlParamName is called when exiting the promqlParamName production.
	ExitPromqlParamName(c *PromqlParamNameContext)

	// ExitPromqlParamValue is called when exiting the promqlParamValue production.
	ExitPromqlParamValue(c *PromqlParamValueContext)

	// ExitPromqlQueryContent is called when exiting the promqlQueryContent production.
	ExitPromqlQueryContent(c *PromqlQueryContentContext)

	// ExitPromqlQueryPart is called when exiting the promqlQueryPart production.
	ExitPromqlQueryPart(c *PromqlQueryPartContext)

	// ExitPromqlIndexPattern is called when exiting the promqlIndexPattern production.
	ExitPromqlIndexPattern(c *PromqlIndexPatternContext)

	// ExitPromqlClusterString is called when exiting the promqlClusterString production.
	ExitPromqlClusterString(c *PromqlClusterStringContext)

	// ExitPromqlSelectorString is called when exiting the promqlSelectorString production.
	ExitPromqlSelectorString(c *PromqlSelectorStringContext)

	// ExitPromqlUnquotedIndexString is called when exiting the promqlUnquotedIndexString production.
	ExitPromqlUnquotedIndexString(c *PromqlUnquotedIndexStringContext)

	// ExitPromqlIndexString is called when exiting the promqlIndexString production.
	ExitPromqlIndexString(c *PromqlIndexStringContext)
}
