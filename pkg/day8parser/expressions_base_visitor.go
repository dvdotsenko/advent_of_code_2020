// Code generated from /mnt/src/Expressions.g4 by ANTLR 4.8. DO NOT EDIT.

package day8parser // Expressions
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseExpressionsVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseExpressionsVisitor) VisitTreeStart(ctx *TreeStartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionsVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}
