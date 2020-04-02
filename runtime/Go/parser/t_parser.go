// Code generated from TParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // TParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)


/* parser/listener/visitor header section */

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 34, 154, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 3, 2, 6, 2, 30, 10, 2, 13, 2, 14, 2, 31, 3, 2, 3, 
	2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 40, 10, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 
	5, 6, 5, 47, 10, 5, 13, 5, 14, 5, 48, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 
	5, 7, 5, 57, 10, 5, 12, 5, 14, 5, 60, 11, 5, 3, 5, 5, 5, 63, 10, 5, 3, 
	5, 5, 5, 66, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 6, 7, 73, 10, 7, 13, 
	7, 14, 7, 74, 3, 7, 5, 7, 78, 10, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 
	3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 90, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 
	3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 101, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 118, 
	10, 9, 12, 9, 14, 9, 121, 11, 9, 3, 10, 3, 10, 3, 10, 5, 10, 126, 10, 10, 
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 134, 10, 12, 12, 12, 14, 
	12, 137, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 145, 
	10, 13, 12, 13, 14, 13, 148, 11, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 
	3, 62, 3, 16, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 2, 4, 
	3, 2, 13, 15, 3, 2, 14, 14, 2, 161, 2, 29, 3, 2, 2, 2, 4, 35, 3, 2, 2, 
	2, 6, 43, 3, 2, 2, 2, 8, 65, 3, 2, 2, 2, 10, 67, 3, 2, 2, 2, 12, 72, 3, 
	2, 2, 2, 14, 89, 3, 2, 2, 2, 16, 100, 3, 2, 2, 2, 18, 125, 3, 2, 2, 2, 
	20, 127, 3, 2, 2, 2, 22, 129, 3, 2, 2, 2, 24, 140, 3, 2, 2, 2, 26, 151, 
	3, 2, 2, 2, 28, 30, 5, 14, 8, 2, 29, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 
	31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 34, 7, 
	2, 2, 3, 34, 3, 3, 2, 2, 2, 35, 39, 7, 8, 2, 2, 36, 37, 5, 6, 4, 2, 37, 
	38, 7, 10, 2, 2, 38, 40, 3, 2, 2, 2, 39, 36, 3, 2, 2, 2, 39, 40, 3, 2, 
	2, 2, 40, 41, 3, 2, 2, 2, 41, 42, 6, 3, 2, 2, 42, 5, 3, 2, 2, 2, 43, 44, 
	7, 12, 2, 2, 44, 7, 3, 2, 2, 2, 45, 47, 5, 4, 3, 2, 46, 45, 3, 2, 2, 2, 
	47, 48, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 66, 3, 
	2, 2, 2, 50, 51, 6, 5, 3, 2, 51, 52, 5, 6, 4, 2, 52, 53, 8, 5, 1, 2, 53, 
	66, 3, 2, 2, 2, 54, 62, 7, 8, 2, 2, 55, 57, 7, 9, 2, 2, 56, 55, 3, 2, 2, 
	2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 61, 
	3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 63, 5, 4, 3, 2, 62, 63, 3, 2, 2, 2, 
	62, 58, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 66, 8, 5, 1, 2, 65, 46, 3, 
	2, 2, 2, 65, 50, 3, 2, 2, 2, 65, 54, 3, 2, 2, 2, 66, 9, 3, 2, 2, 2, 67, 
	68, 5, 14, 8, 2, 68, 11, 3, 2, 2, 2, 69, 70, 5, 10, 6, 2, 70, 71, 11, 2, 
	2, 2, 71, 73, 3, 2, 2, 2, 72, 69, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 72, 
	3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 77, 3, 2, 2, 2, 76, 78, 9, 2, 2, 2, 
	77, 76, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 80, 10, 
	3, 2, 2, 80, 13, 3, 2, 2, 2, 81, 82, 5, 16, 9, 2, 82, 83, 7, 11, 2, 2, 
	83, 84, 5, 16, 9, 2, 84, 85, 7, 14, 2, 2, 85, 90, 3, 2, 2, 2, 86, 87, 5, 
	16, 9, 2, 87, 88, 7, 14, 2, 2, 88, 90, 3, 2, 2, 2, 89, 81, 3, 2, 2, 2, 
	89, 86, 3, 2, 2, 2, 90, 15, 3, 2, 2, 2, 91, 92, 8, 9, 1, 2, 92, 93, 7, 
	18, 2, 2, 93, 94, 5, 16, 9, 2, 94, 95, 7, 19, 2, 2, 95, 101, 3, 2, 2, 2, 
	96, 101, 5, 20, 11, 2, 97, 101, 5, 18, 10, 2, 98, 101, 7, 6, 2, 2, 99, 
	101, 7, 34, 2, 2, 100, 91, 3, 2, 2, 2, 100, 96, 3, 2, 2, 2, 100, 97, 3, 
	2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 99, 3, 2, 2, 2, 101, 119, 3, 2, 2, 2, 
	102, 103, 12, 11, 2, 2, 103, 104, 7, 17, 2, 2, 104, 118, 5, 16, 9, 12, 
	105, 106, 12, 10, 2, 2, 106, 107, 7, 15, 2, 2, 107, 118, 5, 16, 9, 11, 
	108, 109, 12, 8, 2, 2, 109, 110, 7, 22, 2, 2, 110, 111, 5, 16, 9, 2, 111, 
	112, 7, 13, 2, 2, 112, 113, 5, 16, 9, 8, 113, 118, 3, 2, 2, 2, 114, 115, 
	12, 7, 2, 2, 115, 116, 7, 11, 2, 2, 116, 118, 5, 16, 9, 7, 117, 102, 3, 
	2, 2, 2, 117, 105, 3, 2, 2, 2, 117, 108, 3, 2, 2, 2, 117, 114, 3, 2, 2, 
	2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 
	17, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 123, 7, 4, 2, 2, 123, 126, 5, 
	16, 9, 2, 124, 126, 7, 5, 2, 2, 125, 122, 3, 2, 2, 2, 125, 124, 3, 2, 2, 
	2, 126, 19, 3, 2, 2, 2, 127, 128, 7, 8, 2, 2, 128, 21, 3, 2, 2, 2, 129, 
	130, 7, 20, 2, 2, 130, 135, 7, 6, 2, 2, 131, 132, 7, 23, 2, 2, 132, 134, 
	7, 6, 2, 2, 133, 131, 3, 2, 2, 2, 134, 137, 3, 2, 2, 2, 135, 133, 3, 2, 
	2, 2, 135, 136, 3, 2, 2, 2, 136, 138, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 
	138, 139, 7, 21, 2, 2, 139, 23, 3, 2, 2, 2, 140, 141, 7, 20, 2, 2, 141, 
	146, 5, 20, 11, 2, 142, 143, 7, 23, 2, 2, 143, 145, 5, 20, 11, 2, 144, 
	142, 3, 2, 2, 2, 145, 148, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146, 147, 
	3, 2, 2, 2, 147, 149, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 149, 150, 7, 21, 
	2, 2, 150, 25, 3, 2, 2, 2, 151, 152, 11, 2, 2, 2, 152, 27, 3, 2, 2, 2, 
	17, 31, 39, 48, 58, 62, 65, 74, 77, 89, 100, 117, 119, 125, 135, 146,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'return'", "'continue'", "", "", "", "'<'", "'>'", "'='", "'and'", 
	"':'", "';'", "'+'", "'-'", "'*'", "'('", "')'", "'{'", "'}'", "'?'", "','", 
	"", "", "", "", "", "", "'.'", "'..'", "'$'", "'&'",
}
var symbolicNames = []string{
	"", "DUMMY", "Return", "Continue", "INT", "Digit", "ID", "LessThan", "GreaterThan", 
	"Equal", "And", "Colon", "Semicolon", "Plus", "Minus", "Star", "OpenPar", 
	"ClosePar", "OpenCurly", "CloseCurly", "QuestionMark", "Comma", "String", 
	"Foo", "Bar", "Any", "Comment", "WS", "Dot", "DotDot", "Dollar", "Ampersand", 
	"SString",
}

var ruleNames = []string{
	"main", "divide", "and_", "conquer", "unused", "unused2", "stat", "expr", 
	"flowControl", "id", "array", "idarray", "any",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TParser struct {
	*antlr.BaseParser
}

func NewTParser(input antlr.TokenStream) *TParser {
	this := new(TParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "TParser.g4"

	return this
}


/* public parser declarations/members section */
func myAction() bool { return true; }
func doesItBlend() bool { return true; }
func cleanUp() {}
func doInit() {}
func doAfter() {}



// TParser tokens.
const (
	TParserEOF = antlr.TokenEOF
	TParserDUMMY = 1
	TParserReturn = 2
	TParserContinue = 3
	TParserINT = 4
	TParserDigit = 5
	TParserID = 6
	TParserLessThan = 7
	TParserGreaterThan = 8
	TParserEqual = 9
	TParserAnd = 10
	TParserColon = 11
	TParserSemicolon = 12
	TParserPlus = 13
	TParserMinus = 14
	TParserStar = 15
	TParserOpenPar = 16
	TParserClosePar = 17
	TParserOpenCurly = 18
	TParserCloseCurly = 19
	TParserQuestionMark = 20
	TParserComma = 21
	TParserString = 22
	TParserFoo = 23
	TParserBar = 24
	TParserAny = 25
	TParserComment = 26
	TParserWS = 27
	TParserDot = 28
	TParserDotDot = 29
	TParserDollar = 30
	TParserAmpersand = 31
	TParserSString = 32
)

// TParser rules.
const (
	TParserRULE_main = 0
	TParserRULE_divide = 1
	TParserRULE_and_ = 2
	TParserRULE_conquer = 3
	TParserRULE_unused = 4
	TParserRULE_unused2 = 5
	TParserRULE_stat = 6
	TParserRULE_expr = 7
	TParserRULE_flowControl = 8
	TParserRULE_id = 9
	TParserRULE_array = 10
	TParserRULE_idarray = 11
	TParserRULE_any = 12
)

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TParserRULE_main
	return p
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) EOF() antlr.TerminalNode {
	return s.GetToken(TParserEOF, 0)
}

func (s *MainContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *MainContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.ExitMain(s)
	}
}




func (p *TParser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TParserRULE_main)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((((_la - 2)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 2))) & ((1 << (TParserReturn - 2)) | (1 << (TParserContinue - 2)) | (1 << (TParserINT - 2)) | (1 << (TParserID - 2)) | (1 << (TParserOpenPar - 2)) | (1 << (TParserSString - 2)))) != 0) {
		{
			p.SetState(26)
			p.Stat()
		}


		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
		p.Match(TParserEOF)
	}



	return localctx
}


// IDivideContext is an interface to support dynamic dispatch.
type IDivideContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDivideContext differentiates from other interfaces.
	IsDivideContext()
}

type DivideContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDivideContext() *DivideContext {
	var p = new(DivideContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TParserRULE_divide
	return p
}

func (*DivideContext) IsDivideContext() {}

func NewDivideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DivideContext {
	var p = new(DivideContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TParserRULE_divide

	return p
}

func (s *DivideContext) GetParser() antlr.Parser { return s.parser }

func (s *DivideContext) ID() antlr.TerminalNode {
	return s.GetToken(TParserID, 0)
}

func (s *DivideContext) And_() IAnd_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnd_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnd_Context)
}

func (s *DivideContext) GreaterThan() antlr.TerminalNode {
	return s.GetToken(TParserGreaterThan, 0)
}

func (s *DivideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivideContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DivideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.EnterDivide(s)
	}
}

func (s *DivideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.ExitDivide(s)
	}
}




func (p *TParser) Divide() (localctx IDivideContext) {
	localctx = NewDivideContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TParserRULE_divide)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(TParserID)
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(34)
			p.And_()
		}
		{
			p.SetState(35)
			p.Match(TParserGreaterThan)
		}


	}
	p.SetState(39)

	if !(doesItBlend()) {
		panic(antlr.NewFailedPredicateException(p, "doesItBlend()", ""))
	}



	return localctx
}


// IAnd_Context is an interface to support dynamic dispatch.
type IAnd_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnd_Context differentiates from other interfaces.
	IsAnd_Context()
}

type And_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnd_Context() *And_Context {
	var p = new(And_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TParserRULE_and_
	return p
}

func (*And_Context) IsAnd_Context() {}

func NewAnd_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *And_Context {
	var p = new(And_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TParserRULE_and_

	return p
}

func (s *And_Context) GetParser() antlr.Parser { return s.parser }

func (s *And_Context) And() antlr.TerminalNode {
	return s.GetToken(TParserAnd, 0)
}

func (s *And_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *And_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *And_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.EnterAnd_(s)
	}
}

func (s *And_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.ExitAnd_(s)
	}
}




func (p *TParser) And_() (localctx IAnd_Context) {
	localctx = NewAnd_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TParserRULE_and_)
	 doInit(); 

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(TParserAnd)
	}



	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))


	 doAfter(); 


	return localctx
}


// IConquerContext is an interface to support dynamic dispatch.
type IConquerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token 


	// Set_ID sets the _ID token.
	Set_ID(antlr.Token) 


	// IsConquerContext differentiates from other interfaces.
	IsConquerContext()
}

type ConquerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_ID antlr.Token
}

func NewEmptyConquerContext() *ConquerContext {
	var p = new(ConquerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TParserRULE_conquer
	return p
}

func (*ConquerContext) IsConquerContext() {}

func NewConquerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConquerContext {
	var p = new(ConquerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TParserRULE_conquer

	return p
}

func (s *ConquerContext) GetParser() antlr.Parser { return s.parser }

func (s *ConquerContext) Get_ID() antlr.Token { return s._ID }


func (s *ConquerContext) Set_ID(v antlr.Token) { s._ID = v }


func (s *ConquerContext) AllDivide() []IDivideContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDivideContext)(nil)).Elem())
	var tst = make([]IDivideContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDivideContext)
		}
	}

	return tst
}

func (s *ConquerContext) Divide(i int) IDivideContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDivideContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDivideContext)
}

func (s *ConquerContext) And_() IAnd_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnd_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnd_Context)
}

func (s *ConquerContext) ID() antlr.TerminalNode {
	return s.GetToken(TParserID, 0)
}

func (s *ConquerContext) AllLessThan() []antlr.TerminalNode {
	return s.GetTokens(TParserLessThan)
}

func (s *ConquerContext) LessThan(i int) antlr.TerminalNode {
	return s.GetToken(TParserLessThan, i)
}

func (s *ConquerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConquerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConquerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.EnterConquer(s)
	}
}

func (s *ConquerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.ExitConquer(s)
	}
}




func (p *TParser) Conquer() (localctx IConquerContext) {
	localctx = NewConquerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TParserRULE_conquer)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == TParserID {
			{
				p.SetState(43)
				p.Divide()
			}


			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(48)

		if !(doesItBlend()) {
			panic(antlr.NewFailedPredicateException(p, "doesItBlend()", ""))
		}
		{
			p.SetState(49)
			p.And_()
		}
		 myAction(); 


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(52)

			var _m = p.Match(TParserID)

			localctx.(*ConquerContext)._ID = _m
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1+1 {
			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			for _la == TParserLessThan {
				{
					p.SetState(53)
					p.Match(TParserLessThan)
				}


				p.SetState(58)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(59)
				p.Divide()
			}


		}
		 (func() string { if localctx.(*ConquerContext).Get_ID() == nil { return "" } else { return localctx.(*ConquerContext).Get_ID().GetText() }}()); 

	}


	return localctx
}


// IUnusedContext is an interface to support dynamic dispatch.
type IUnusedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInput returns the input attribute.
	GetInput() float64

	// GetCalculated returns the calculated attribute.
	GetCalculated() float64

	// Get_a returns the _a attribute.
	Get_a() int

	// Get_b returns the _b attribute.
	Get_b() float64

	// Get_c returns the _c attribute.
	Get_c() int


	// SetInput sets the input attribute.
	SetInput(float64)

	// SetCalculated sets the calculated attribute.
	SetCalculated(float64)

	// Set_a sets the _a attribute.
	Set_a(int)

	// Set_b sets the _b attribute.
	Set_b(float64)

	// Set_c sets the _c attribute.
	Set_c(int)


	// IsUnusedContext differentiates from other interfaces.
	IsUnusedContext()
}

type UnusedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	input float64// TODO = 111
	calculated float64
	_a int
	_b float64
	_c int
}

func NewEmptyUnusedContext() *UnusedContext {
	var p = new(UnusedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TParserRULE_unused
	return p
}

func (*UnusedContext) IsUnusedContext() {}

func NewUnusedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, input float64) *UnusedContext {
	var p = new(UnusedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TParserRULE_unused

	p.input = input

	return p
}

func (s *UnusedContext) GetParser() antlr.Parser { return s.parser }

func (s *UnusedContext) GetInput() float64 { return s.input }

func (s *UnusedContext) GetCalculated() float64 { return s.calculated }

func (s *UnusedContext) Get_a() int { return s._a }

func (s *UnusedContext) Get_b() float64 { return s._b }

func (s *UnusedContext) Get_c() int { return s._c }


func (s *UnusedContext) SetInput(v float64) { s.input = v }

func (s *UnusedContext) SetCalculated(v float64) { s.calculated = v }

func (s *UnusedContext) Set_a(v int) { s._a = v }

func (s *UnusedContext) Set_b(v float64) { s._b = v }

func (s *UnusedContext) Set_c(v int) { s._c = v }


func (s *UnusedContext) Stat() IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *UnusedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnusedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UnusedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.EnterUnused(s)
	}
}

func (s *UnusedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.ExitUnused(s)
	}
}




func (p *TParser) Unused(input float64) (localctx IUnusedContext) {
	localctx = NewUnusedContext(p, p.GetParserRuleContext(), p.GetState(), input)
	p.EnterRule(localctx, 8, TParserRULE_unused)
	 doInit(); 

	defer func() {

		  cleanUp();

		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Stat()
	}



	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))


	 doAfter(); 


	return localctx
}


// IUnused2Context is an interface to support dynamic dispatch.
type IUnused2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnused2Context differentiates from other interfaces.
	IsUnused2Context()
}

type Unused2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnused2Context() *Unused2Context {
	var p = new(Unused2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TParserRULE_unused2
	return p
}

func (*Unused2Context) IsUnused2Context() {}

func NewUnused2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unused2Context {
	var p = new(Unused2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TParserRULE_unused2

	return p
}

func (s *Unused2Context) GetParser() antlr.Parser { return s.parser }

func (s *Unused2Context) AllSemicolon() []antlr.TerminalNode {
	return s.GetTokens(TParserSemicolon)
}

func (s *Unused2Context) Semicolon(i int) antlr.TerminalNode {
	return s.GetToken(TParserSemicolon, i)
}

func (s *Unused2Context) AllUnused() []IUnusedContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnusedContext)(nil)).Elem())
	var tst = make([]IUnusedContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnusedContext)
		}
	}

	return tst
}

func (s *Unused2Context) Unused(i int) IUnusedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnusedContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnusedContext)
}

func (s *Unused2Context) Colon() antlr.TerminalNode {
	return s.GetToken(TParserColon, 0)
}

func (s *Unused2Context) Plus() antlr.TerminalNode {
	return s.GetToken(TParserPlus, 0)
}

func (s *Unused2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unused2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Unused2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.EnterUnused2(s)
	}
}

func (s *Unused2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.ExitUnused2(s)
	}
}




func (p *TParser) Unused2() (localctx IUnused2Context) {
	localctx = NewUnused2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TParserRULE_unused2)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(67)
					p.Unused(1)
				}
				p.SetState(68)
				p.MatchWildcard()





		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(74)
			_la = p.GetTokenStream().LA(1)

			if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << TParserColon) | (1 << TParserSemicolon) | (1 << TParserPlus))) != 0)) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}


	}
	{
		p.SetState(77)
		_la = p.GetTokenStream().LA(1)

		if _la <= 0 || _la == TParserSemicolon  {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *StatContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StatContext) Equal() antlr.TerminalNode {
	return s.GetToken(TParserEqual, 0)
}

func (s *StatContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(TParserSemicolon, 0)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.ExitStat(s)
	}
}




func (p *TParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TParserRULE_stat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.expr(0)
		}
		{
			p.SetState(80)
			p.Match(TParserEqual)
		}
		{
			p.SetState(81)
			p.expr(0)
		}
		{
			p.SetState(82)
			p.Match(TParserSemicolon)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.expr(0)
		}
		{
			p.SetState(85)
			p.Match(TParserSemicolon)
		}

	}


	return localctx
}


// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdentifier returns the identifier rule contexts.
	GetIdentifier() IIdContext


	// SetIdentifier sets the identifier rule contexts.
	SetIdentifier(IIdContext)


	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	identifier IIdContext 
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetIdentifier() IIdContext { return s.identifier }


func (s *ExprContext) SetIdentifier(v IIdContext) { s.identifier = v }


func (s *ExprContext) OpenPar() antlr.TerminalNode {
	return s.GetToken(TParserOpenPar, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) ClosePar() antlr.TerminalNode {
	return s.GetToken(TParserClosePar, 0)
}

func (s *ExprContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *ExprContext) FlowControl() IFlowControlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFlowControlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFlowControlContext)
}

func (s *ExprContext) INT() antlr.TerminalNode {
	return s.GetToken(TParserINT, 0)
}

func (s *ExprContext) SString() antlr.TerminalNode {
	return s.GetToken(TParserSString, 0)
}

func (s *ExprContext) Star() antlr.TerminalNode {
	return s.GetToken(TParserStar, 0)
}

func (s *ExprContext) Plus() antlr.TerminalNode {
	return s.GetToken(TParserPlus, 0)
}

func (s *ExprContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(TParserQuestionMark, 0)
}

func (s *ExprContext) Colon() antlr.TerminalNode {
	return s.GetToken(TParserColon, 0)
}

func (s *ExprContext) Equal() antlr.TerminalNode {
	return s.GetToken(TParserEqual, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.ExitExpr(s)
	}
}





func (p *TParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *TParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, TParserRULE_expr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TParserOpenPar:
		{
			p.SetState(90)
			p.Match(TParserOpenPar)
		}
		{
			p.SetState(91)
			p.expr(0)
		}
		{
			p.SetState(92)
			p.Match(TParserClosePar)
		}


	case TParserID:
		{
			p.SetState(94)

			var _x = p.Id()


			localctx.(*ExprContext).identifier = _x
		}


	case TParserReturn, TParserContinue:
		{
			p.SetState(95)
			p.FlowControl()
		}


	case TParserINT:
		{
			p.SetState(96)
			p.Match(TParserINT)
		}


	case TParserSString:
		{
			p.SetState(97)
			p.Match(TParserSString)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(115)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TParserRULE_expr)
				p.SetState(100)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(101)
					p.Match(TParserStar)
				}
				{
					p.SetState(102)
					p.expr(10)
				}


			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TParserRULE_expr)
				p.SetState(103)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(104)
					p.Match(TParserPlus)
				}
				{
					p.SetState(105)
					p.expr(9)
				}


			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TParserRULE_expr)
				p.SetState(106)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(107)
					p.Match(TParserQuestionMark)
				}
				{
					p.SetState(108)
					p.expr(0)
				}
				{
					p.SetState(109)
					p.Match(TParserColon)
				}
				{
					p.SetState(110)
					p.expr(6)
				}


			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TParserRULE_expr)
				p.SetState(112)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(113)
					p.Match(TParserEqual)
				}
				{
					p.SetState(114)
					p.expr(5)
				}

			}

		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}



	return localctx
}


// IFlowControlContext is an interface to support dynamic dispatch.
type IFlowControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFlowControlContext differentiates from other interfaces.
	IsFlowControlContext()
}

type FlowControlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlowControlContext() *FlowControlContext {
	var p = new(FlowControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TParserRULE_flowControl
	return p
}

func (*FlowControlContext) IsFlowControlContext() {}

func NewFlowControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlowControlContext {
	var p = new(FlowControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TParserRULE_flowControl

	return p
}

func (s *FlowControlContext) GetParser() antlr.Parser { return s.parser }

func (s *FlowControlContext) CopyFrom(ctx *FlowControlContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FlowControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlowControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type ReturnContext struct {
	*FlowControlContext
}

func NewReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnContext {
	var p = new(ReturnContext)

	p.FlowControlContext = NewEmptyFlowControlContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FlowControlContext))

	return p
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) Return() antlr.TerminalNode {
	return s.GetToken(TParserReturn, 0)
}

func (s *ReturnContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.ExitReturn(s)
	}
}


type ContinueContext struct {
	*FlowControlContext
}

func NewContinueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueContext {
	var p = new(ContinueContext)

	p.FlowControlContext = NewEmptyFlowControlContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FlowControlContext))

	return p
}

func (s *ContinueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueContext) Continue() antlr.TerminalNode {
	return s.GetToken(TParserContinue, 0)
}


func (s *ContinueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.EnterContinue(s)
	}
}

func (s *ContinueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.ExitContinue(s)
	}
}



func (p *TParser) FlowControl() (localctx IFlowControlContext) {
	localctx = NewFlowControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TParserRULE_flowControl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TParserReturn:
		localctx = NewReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Match(TParserReturn)
		}
		{
			p.SetState(121)
			p.expr(0)
		}


	case TParserContinue:
		localctx = NewContinueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Match(TParserContinue)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IIdContext is an interface to support dynamic dispatch.
type IIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdContext differentiates from other interfaces.
	IsIdContext()
}

type IdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdContext() *IdContext {
	var p = new(IdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TParserRULE_id
	return p
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(TParserID, 0)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.ExitId(s)
	}
}




func (p *TParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TParserRULE_id)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(TParserID)
	}



	return localctx
}


// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_INT returns the _INT token.
	Get_INT() antlr.Token 


	// Set_INT sets the _INT token.
	Set_INT(antlr.Token) 


	// GetEl returns the el token list.
	GetEl() []antlr.Token


	// SetEl sets the el token list.
	SetEl([]antlr.Token)


	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_INT antlr.Token
	el []antlr.Token
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) Get_INT() antlr.Token { return s._INT }


func (s *ArrayContext) Set_INT(v antlr.Token) { s._INT = v }


func (s *ArrayContext) GetEl() []antlr.Token { return s.el }


func (s *ArrayContext) SetEl(v []antlr.Token) { s.el = v }


func (s *ArrayContext) OpenCurly() antlr.TerminalNode {
	return s.GetToken(TParserOpenCurly, 0)
}

func (s *ArrayContext) CloseCurly() antlr.TerminalNode {
	return s.GetToken(TParserCloseCurly, 0)
}

func (s *ArrayContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(TParserINT)
}

func (s *ArrayContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(TParserINT, i)
}

func (s *ArrayContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(TParserComma)
}

func (s *ArrayContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(TParserComma, i)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.ExitArray(s)
	}
}




func (p *TParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TParserRULE_array)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(TParserOpenCurly)
	}
	{
		p.SetState(128)

		var _m = p.Match(TParserINT)

		localctx.(*ArrayContext)._INT = _m
	}
	localctx.(*ArrayContext).el = append(localctx.(*ArrayContext).el, localctx.(*ArrayContext)._INT)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == TParserComma {
		{
			p.SetState(129)
			p.Match(TParserComma)
		}
		{
			p.SetState(130)

			var _m = p.Match(TParserINT)

			localctx.(*ArrayContext)._INT = _m
		}
		localctx.(*ArrayContext).el = append(localctx.(*ArrayContext).el, localctx.(*ArrayContext)._INT)


		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(136)
		p.Match(TParserCloseCurly)
	}



	return localctx
}


// IIdarrayContext is an interface to support dynamic dispatch.
type IIdarrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_id returns the _id rule contexts.
	Get_id() IIdContext


	// Set_id sets the _id rule contexts.
	Set_id(IIdContext)


	// GetElement returns the element rule context list.
	GetElement() []IIdContext


	// SetElement sets the element rule context list.
	SetElement([]IIdContext) 


	// IsIdarrayContext differentiates from other interfaces.
	IsIdarrayContext()
}

type IdarrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_id IIdContext 
	element []IIdContext
}

func NewEmptyIdarrayContext() *IdarrayContext {
	var p = new(IdarrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TParserRULE_idarray
	return p
}

func (*IdarrayContext) IsIdarrayContext() {}

func NewIdarrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdarrayContext {
	var p = new(IdarrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TParserRULE_idarray

	return p
}

func (s *IdarrayContext) GetParser() antlr.Parser { return s.parser }

func (s *IdarrayContext) Get_id() IIdContext { return s._id }


func (s *IdarrayContext) Set_id(v IIdContext) { s._id = v }


func (s *IdarrayContext) GetElement() []IIdContext { return s.element }


func (s *IdarrayContext) SetElement(v []IIdContext) { s.element = v }


func (s *IdarrayContext) OpenCurly() antlr.TerminalNode {
	return s.GetToken(TParserOpenCurly, 0)
}

func (s *IdarrayContext) CloseCurly() antlr.TerminalNode {
	return s.GetToken(TParserCloseCurly, 0)
}

func (s *IdarrayContext) AllId() []IIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdContext)(nil)).Elem())
	var tst = make([]IIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdContext)
		}
	}

	return tst
}

func (s *IdarrayContext) Id(i int) IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *IdarrayContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(TParserComma)
}

func (s *IdarrayContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(TParserComma, i)
}

func (s *IdarrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdarrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IdarrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.EnterIdarray(s)
	}
}

func (s *IdarrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.ExitIdarray(s)
	}
}




func (p *TParser) Idarray() (localctx IIdarrayContext) {
	localctx = NewIdarrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TParserRULE_idarray)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(TParserOpenCurly)
	}
	{
		p.SetState(139)

		var _x = p.Id()


		localctx.(*IdarrayContext)._id = _x
	}
	localctx.(*IdarrayContext).element = append(localctx.(*IdarrayContext).element, localctx.(*IdarrayContext)._id)
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == TParserComma {
		{
			p.SetState(140)
			p.Match(TParserComma)
		}
		{
			p.SetState(141)

			var _x = p.Id()


			localctx.(*IdarrayContext)._id = _x
		}
		localctx.(*IdarrayContext).element = append(localctx.(*IdarrayContext).element, localctx.(*IdarrayContext)._id)


		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(147)
		p.Match(TParserCloseCurly)
	}



	return localctx
}


// IAnyContext is an interface to support dynamic dispatch.
type IAnyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token 


	// SetT sets the t token.
	SetT(antlr.Token) 


	// IsAnyContext differentiates from other interfaces.
	IsAnyContext()
}

type AnyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	t antlr.Token
}

func NewEmptyAnyContext() *AnyContext {
	var p = new(AnyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TParserRULE_any
	return p
}

func (*AnyContext) IsAnyContext() {}

func NewAnyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnyContext {
	var p = new(AnyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TParserRULE_any

	return p
}

func (s *AnyContext) GetParser() antlr.Parser { return s.parser }

func (s *AnyContext) GetT() antlr.Token { return s.t }


func (s *AnyContext) SetT(v antlr.Token) { s.t = v }

func (s *AnyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AnyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.EnterAny(s)
	}
}

func (s *AnyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TParserListener); ok {
		listenerT.ExitAny(s)
	}
}




func (p *TParser) Any() (localctx IAnyContext) {
	localctx = NewAnyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TParserRULE_any)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(149)

	var _mwc = p.MatchWildcard()

	localctx.(*AnyContext).t = _mwc




	return localctx
}


func (p *TParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
			var t *DivideContext = nil
			if localctx != nil { t = localctx.(*DivideContext) }
			return p.Divide_Sempred(t, predIndex)

	case 3:
			var t *ConquerContext = nil
			if localctx != nil { t = localctx.(*ConquerContext) }
			return p.Conquer_Sempred(t, predIndex)

	case 7:
			var t *ExprContext = nil
			if localctx != nil { t = localctx.(*ExprContext) }
			return p.Expr_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TParser) Divide_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return doesItBlend()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TParser) Conquer_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
			return doesItBlend()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
			return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
			return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
			return p.Precpred(p.GetParserRuleContext(), 6)

	case 5:
			return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

