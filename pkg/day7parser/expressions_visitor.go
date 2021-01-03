// Code generated from /mnt/src/Expressions.g4 by ANTLR 4.8. DO NOT EDIT.

package day7parser // Expressions
import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by ExpressionsParser.
type ExpressionsVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExpressionsParser#treeStart.
	VisitTreeStart(ctx *TreeStartContext) interface{}

	// Visit a parse tree produced by ExpressionsParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by ExpressionsParser#component.
	VisitComponent(ctx *ComponentContext) interface{}

	// Visit a parse tree produced by ExpressionsParser#label.
	VisitLabel(ctx *LabelContext) interface{}

}