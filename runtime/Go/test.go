package main

import (
	"fmt"
	"os"

	parser "./cppparser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type treeShapeListener struct {
	*parser.BaseCPP14Listener
}

func newTreeShapeListener() *treeShapeListener {
	return new(treeShapeListener)
}

func (l *treeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
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
}
