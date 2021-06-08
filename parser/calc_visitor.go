// Code generated from /home/hsc/GolandProjects/go/src/antr4_example/Calc.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CalcParser.
type CalcVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalcParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by CalcParser#Parenthesis.
	VisitParenthesis(ctx *ParenthesisContext) interface{}

	// Visit a parse tree produced by CalcParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by CalcParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by CalcParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}
}
