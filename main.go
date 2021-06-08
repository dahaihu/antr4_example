package main

import (
	"antr4_example/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)


func main() {
	// Setup the input
	is := antlr.NewInputStream("1 + 2 * 3")
	//fmt.Println(calc("1+2+(3+4)*(5-10)"))
	//
	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)
	tree := p.Start()

	visitor := CalcVisitor{}
	fmt.Println(visitor.Visit(tree))
}