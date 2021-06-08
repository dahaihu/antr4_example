package main

import (
	"antr4_example/parser"
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CalcVisitor struct {
	parser.BaseCalcVisitor
}

func (v *CalcVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.AddSubContext:
		return v.VisitAddSub(val)
	case *parser.MulDivContext:
		return v.VisitMulDiv(val)
	case *parser.NumberContext:
		return v.VisitNumber(val)
	case *parser.StartContext:
		return v.Visit(val.Expression())
	case *parser.ParenthesisContext:
		fmt.Println(tree.GetChildren())
		return nil
	default:
		panic("Unknown context")
	}
}

func (v *CalcVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	left := v.Visit(ctx.Expression(0)).(int)
	right := v.Visit(ctx.Expression(1)).(int)

	switch ctx.GetOp().GetTokenType() {
	case parser.CalcParserADD:
		return left+right
	case parser.CalcParserSUB:
		return left-right
	default:
		panic(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
	}
}

func (v *CalcVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	left := v.Visit(ctx.Expression(0)).(int)
	right := v.Visit(ctx.Expression(1)).(int)
	switch ctx.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		return left * right
	case parser.CalcParserDIV:
		return left / right
	default:
		panic(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
	}
}

func (v *CalcVisitor) VisitNumber(ctx *parser.NumberContext) interface{} {
	val, _ := strconv.Atoi(ctx.NUMBER().GetText())
	return val
}
