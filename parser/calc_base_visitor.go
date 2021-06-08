// Code generated from /home/hsc/GolandProjects/go/src/antr4_example/Calc.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Calc

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"fmt"
)

type BaseCalcVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCalcVisitor) VisitStart(ctx *StartContext) interface{} {
	fmt.Println("start is ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *BaseCalcVisitor) VisitParenthesis(ctx *ParenthesisContext) interface{} {
	fmt.Println("parenthesis is ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *BaseCalcVisitor) VisitNumber(ctx *NumberContext) interface{} {
	fmt.Println("number is ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *BaseCalcVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	fmt.Println("mul div is ", ctx.GetText())
	return v.VisitChildren(ctx)
}

func (v *BaseCalcVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	fmt.Println("add sub is ", ctx.GetText())
	return v.VisitChildren(ctx)
}
