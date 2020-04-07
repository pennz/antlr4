package main

import (
	"fmt"
	"os"

	parser "./cppparser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type SourceChecker struct {
	p *parser.CPP14Parser
	*parser.CPP14Lexer
	markerList []string
}

type treeShapeListener struct {
	*parser.BaseCPP14Listener
}

func newTreeShapeListener() *treeShapeListener {
	return new(treeShapeListener)
}

func (l *treeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("ignore", ctx)
}

func (c *SourceChecker) init(input antlr.CharStream) {
	c.CPP14Lexer = parser.NewCPP14Lexer(input)
	stream := antlr.NewCommonTokenStream(c.CPP14Lexer, 0)
	c.p = parser.NewCPP14Parser(stream)
}

func (c *SourceChecker) GetMarkerList() []string {
	return make([]string, 0)
}

func (c *SourceChecker) SetMarkerList(fileName string) {

}

type fileInfo struct {
	name string
	info []infoItem
}
type infoItem struct {
	name     string
	location string
}

func (c *SourceChecker) CheckFile() fileInfo {
	return fileInfo{"", nil}
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewCPP14Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCPP14Parser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Translationunit()
	antlr.ParseTreeWalkerDefault.Walk(newTreeShapeListener(), tree)
	fmt.Println(tree)
}
