// Code generated from TLexer.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

/* lexer header section */

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 33, 211, 
	8, 1, 8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 
	9, 6, 4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 
	12, 9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 
	9, 17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 
	22, 4, 23, 9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 
	4, 28, 9, 28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 
	33, 9, 33, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 6, 4, 87, 10, 4, 13, 4, 14, 4, 
	88, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 7, 6, 96, 10, 6, 12, 6, 14, 6, 99, 11, 
	6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 
	3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 
	16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 
	3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 
	23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 7, 25, 152, 
	10, 25, 12, 25, 14, 25, 155, 11, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 
	27, 3, 27, 3, 28, 3, 28, 3, 28, 5, 28, 177, 10, 28, 3, 28, 3, 28, 3, 28, 
	3, 29, 3, 29, 7, 29, 184, 10, 29, 12, 29, 14, 29, 187, 11, 29, 3, 29, 5, 
	29, 190, 10, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 6, 30, 197, 10, 30, 
	13, 30, 14, 30, 198, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 
	32, 3, 33, 3, 33, 3, 33, 3, 153, 2, 34, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 
	15, 2, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 
	17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 32, 49, 33, 51, 
	24, 53, 25, 55, 26, 57, 27, 59, 28, 61, 29, 63, 2, 65, 30, 67, 31, 5, 2, 
	3, 4, 5, 3, 2, 50, 59, 4, 2, 12, 12, 15, 15, 5, 2, 11, 12, 15, 15, 34, 
	34, 3, 5, 2, 67, 2, 92, 2, 99, 2, 124, 2, 130, 2, 1, 18, 214, 2, 5, 3, 
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 
	2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 
	2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 
	2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 
	2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 
	3, 2, 2, 2, 3, 65, 3, 2, 2, 2, 4, 67, 3, 2, 2, 2, 5, 69, 3, 2, 2, 2, 7, 
	76, 3, 2, 2, 2, 9, 86, 3, 2, 2, 2, 11, 90, 3, 2, 2, 2, 13, 92, 3, 2, 2, 
	2, 15, 100, 3, 2, 2, 2, 17, 102, 3, 2, 2, 2, 19, 104, 3, 2, 2, 2, 21, 106, 
	3, 2, 2, 2, 23, 108, 3, 2, 2, 2, 25, 112, 3, 2, 2, 2, 27, 114, 3, 2, 2, 
	2, 29, 116, 3, 2, 2, 2, 31, 118, 3, 2, 2, 2, 33, 120, 3, 2, 2, 2, 35, 122, 
	3, 2, 2, 2, 37, 124, 3, 2, 2, 2, 39, 126, 3, 2, 2, 2, 41, 130, 3, 2, 2, 
	2, 43, 134, 3, 2, 2, 2, 45, 136, 3, 2, 2, 2, 47, 140, 3, 2, 2, 2, 49, 145, 
	3, 2, 2, 2, 51, 149, 3, 2, 2, 2, 53, 158, 3, 2, 2, 2, 55, 166, 3, 2, 2, 
	2, 57, 173, 3, 2, 2, 2, 59, 181, 3, 2, 2, 2, 61, 196, 3, 2, 2, 2, 63, 202, 
	3, 2, 2, 2, 65, 206, 3, 2, 2, 2, 67, 208, 3, 2, 2, 2, 69, 70, 7, 116, 2, 
	2, 70, 71, 7, 103, 2, 2, 71, 72, 7, 118, 2, 2, 72, 73, 7, 119, 2, 2, 73, 
	74, 7, 116, 2, 2, 74, 75, 7, 112, 2, 2, 75, 6, 3, 2, 2, 2, 76, 77, 7, 101, 
	2, 2, 77, 78, 7, 113, 2, 2, 78, 79, 7, 112, 2, 2, 79, 80, 7, 118, 2, 2, 
	80, 81, 7, 107, 2, 2, 81, 82, 7, 112, 2, 2, 82, 83, 7, 119, 2, 2, 83, 84, 
	7, 103, 2, 2, 84, 8, 3, 2, 2, 2, 85, 87, 5, 11, 5, 2, 86, 85, 3, 2, 2, 
	2, 87, 88, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 10, 
	3, 2, 2, 2, 90, 91, 9, 2, 2, 2, 91, 12, 3, 2, 2, 2, 92, 97, 5, 15, 7, 2, 
	93, 96, 5, 15, 7, 2, 94, 96, 4, 50, 59, 2, 95, 93, 3, 2, 2, 2, 95, 94, 
	3, 2, 2, 2, 96, 99, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 
	98, 14, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 100, 101, 9, 5, 2, 2, 101, 16, 
	3, 2, 2, 2, 102, 103, 7, 62, 2, 2, 103, 18, 3, 2, 2, 2, 104, 105, 7, 64, 
	2, 2, 105, 20, 3, 2, 2, 2, 106, 107, 7, 63, 2, 2, 107, 22, 3, 2, 2, 2, 
	108, 109, 7, 99, 2, 2, 109, 110, 7, 112, 2, 2, 110, 111, 7, 102, 2, 2, 
	111, 24, 3, 2, 2, 2, 112, 113, 7, 60, 2, 2, 113, 26, 3, 2, 2, 2, 114, 115, 
	7, 61, 2, 2, 115, 28, 3, 2, 2, 2, 116, 117, 7, 45, 2, 2, 117, 30, 3, 2, 
	2, 2, 118, 119, 7, 47, 2, 2, 119, 32, 3, 2, 2, 2, 120, 121, 7, 44, 2, 2, 
	121, 34, 3, 2, 2, 2, 122, 123, 7, 42, 2, 2, 123, 36, 3, 2, 2, 2, 124, 125, 
	7, 43, 2, 2, 125, 38, 3, 2, 2, 2, 126, 127, 7, 125, 2, 2, 127, 128, 3, 
	2, 2, 2, 128, 129, 8, 19, 2, 2, 129, 40, 3, 2, 2, 2, 130, 131, 7, 127, 
	2, 2, 131, 132, 3, 2, 2, 2, 132, 133, 8, 20, 3, 2, 133, 42, 3, 2, 2, 2, 
	134, 135, 7, 65, 2, 2, 135, 44, 3, 2, 2, 2, 136, 137, 7, 46, 2, 2, 137, 
	138, 3, 2, 2, 2, 138, 139, 8, 22, 4, 2, 139, 46, 3, 2, 2, 2, 140, 141, 
	7, 38, 2, 2, 141, 142, 3, 2, 2, 2, 142, 143, 8, 23, 5, 2, 143, 144, 8, 
	23, 6, 2, 144, 48, 3, 2, 2, 2, 145, 146, 7, 40, 2, 2, 146, 147, 3, 2, 2, 
	2, 147, 148, 8, 24, 7, 2, 148, 50, 3, 2, 2, 2, 149, 153, 7, 36, 2, 2, 150, 
	152, 11, 2, 2, 2, 151, 150, 3, 2, 2, 2, 152, 155, 3, 2, 2, 2, 153, 154, 
	3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 154, 156, 3, 2, 2, 2, 155, 153, 3, 2, 
	2, 2, 156, 157, 7, 36, 2, 2, 157, 52, 3, 2, 2, 2, 158, 159, 6, 26, 2, 2, 
	159, 160, 7, 104, 2, 2, 160, 161, 7, 113, 2, 2, 161, 162, 7, 113, 2, 2, 
	162, 163, 3, 2, 2, 2, 163, 164, 6, 26, 3, 2, 164, 165, 8, 26, 8, 2, 165, 
	54, 3, 2, 2, 2, 166, 167, 7, 100, 2, 2, 167, 168, 7, 99, 2, 2, 168, 169, 
	7, 116, 2, 2, 169, 170, 3, 2, 2, 2, 170, 171, 6, 27, 4, 2, 171, 172, 8, 
	27, 9, 2, 172, 56, 3, 2, 2, 2, 173, 174, 5, 53, 26, 2, 174, 176, 5, 65, 
	32, 2, 175, 177, 5, 55, 27, 2, 176, 175, 3, 2, 2, 2, 176, 177, 3, 2, 2, 
	2, 177, 178, 3, 2, 2, 2, 178, 179, 5, 67, 33, 2, 179, 180, 5, 63, 31, 2, 
	180, 58, 3, 2, 2, 2, 181, 185, 7, 37, 2, 2, 182, 184, 10, 3, 2, 2, 183, 
	182, 3, 2, 2, 2, 184, 187, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 185, 186, 
	3, 2, 2, 2, 186, 189, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 188, 190, 7, 15, 
	2, 2, 189, 188, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 
	191, 192, 7, 12, 2, 2, 192, 193, 3, 2, 2, 2, 193, 194, 8, 29, 10, 2, 194, 
	60, 3, 2, 2, 2, 195, 197, 9, 4, 2, 2, 196, 195, 3, 2, 2, 2, 197, 198, 3, 
	2, 2, 2, 198, 196, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 200, 3, 2, 2, 
	2, 200, 201, 8, 30, 11, 2, 201, 62, 3, 2, 2, 2, 202, 203, 7, 68, 2, 2, 
	203, 204, 7, 99, 2, 2, 204, 205, 7, 124, 2, 2, 205, 64, 3, 2, 2, 2, 206, 
	207, 7, 48, 2, 2, 207, 66, 3, 2, 2, 2, 208, 209, 7, 48, 2, 2, 209, 210, 
	7, 48, 2, 2, 210, 68, 3, 2, 2, 2, 13, 2, 3, 4, 88, 95, 97, 153, 176, 185, 
	189, 198, 12, 7, 3, 2, 6, 2, 2, 8, 2, 2, 5, 2, 2, 4, 3, 2, 9, 3, 2, 3, 
	26, 2, 3, 27, 3, 2, 4, 2, 2, 101, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN", "CommentsChannel", "DirectiveChannel",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "Mode1", "Mode2",
}

var lexerLiteralNames = []string{
	"", "", "'return'", "'continue'", "", "", "", "'<'", "'>'", "'='", "'and'", 
	"':'", "';'", "'+'", "'-'", "'*'", "'('", "')'", "'{'", "'}'", "'?'", "','", 
	"", "", "", "", "", "", "'.'", "'..'", "'$'", "'&'",
}

var lexerSymbolicNames = []string{
	"", "DUMMY", "Return", "Continue", "INT", "Digit", "ID", "LessThan", "GreaterThan", 
	"Equal", "And", "Colon", "Semicolon", "Plus", "Minus", "Star", "OpenPar", 
	"ClosePar", "OpenCurly", "CloseCurly", "QuestionMark", "Comma", "SString", 
	"Foo", "Bar", "Any", "Comment", "WS", "Dot", "DotDot", "Dollar", "Ampersand",
}

var lexerRuleNames = []string{
	"Return", "Continue", "INT", "Digit", "ID", "LETTER", "LessThan", "GreaterThan", 
	"Equal", "And", "Colon", "Semicolon", "Plus", "Minus", "Star", "OpenPar", 
	"ClosePar", "OpenCurly", "CloseCurly", "QuestionMark", "Comma", "Dollar", 
	"Ampersand", "SString", "Foo", "Bar", "Any", "Comment", "WS", "Baz", "Dot", 
	"DotDot",
}

type TLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewTLexer(input antlr.CharStream) *TLexer {

	l := new(TLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "TLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TLexer tokens.
const (
	TLexerDUMMY = 1
	TLexerReturn = 2
	TLexerContinue = 3
	TLexerINT = 4
	TLexerDigit = 5
	TLexerID = 6
	TLexerLessThan = 7
	TLexerGreaterThan = 8
	TLexerEqual = 9
	TLexerAnd = 10
	TLexerColon = 11
	TLexerSemicolon = 12
	TLexerPlus = 13
	TLexerMinus = 14
	TLexerStar = 15
	TLexerOpenPar = 16
	TLexerClosePar = 17
	TLexerOpenCurly = 18
	TLexerCloseCurly = 19
	TLexerQuestionMark = 20
	TLexerComma = 21
	TLexerSString = 22
	TLexerFoo = 23
	TLexerBar = 24
	TLexerAny = 25
	TLexerComment = 26
	TLexerWS = 27
	TLexerDot = 28
	TLexerDotDot = 29
	TLexerDollar = 30
	TLexerAmpersand = 31
)

// TLexer channels.
const (
	TLexerCommentsChannel = 2
	TLexerDirectiveChannel = 3
)

// TLexer modes.
const (
	TLexerMode1 = iota + 1
	TLexerMode2
)

/* public lexer declarations section */
func canTestFoo() bool { return true; }
func isItFoo() bool { return true; }
func isItBar() bool { return true; }

func myFooLexerAction() { /* do something*/ };
func myBarLexerAction() { /* do something*/ };



func (l *TLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 24:
			l.Foo_Action(localctx, actionIndex)


	case 25:
			l.Bar_Action(localctx, actionIndex)


	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *TLexer) Foo_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 0:
			 myFooLexerAction(); 


	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *TLexer) Bar_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 1:
			 myBarLexerAction(); 


	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}


func (l *TLexer) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 24:
			return l.Foo_Sempred(localctx, predIndex)


	case 25:
			return l.Bar_Sempred(localctx, predIndex)



	default:
		panic("No registered predicate for: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TLexer) Foo_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return canTestFoo()

	case 1:
			return isItFoo()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TLexer) Bar_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
			return isItBar()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
