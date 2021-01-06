package day8parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

type Visitor struct {
	// This embed is redundant and is used only to express "inheritance" tree
	// All methods from this and its parent must be reimplemented here at top level
	// because almost all of them are NOOP by default
	*BaseExpressionsVisitor
}

// -----------------------------
// Override BaseParseTreeVisitor

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	//fmt.Printf("Visit : %s '%v'\n", introspection.GetTypeName(tree), tree.(antlr.ParseTree).GetText())
	return tree.Accept(v)
}

//func (v *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
//	rr := make([]interface{}, 2)
//	for _, n := range node.GetChildren() {
//		fmt.Printf("Child : %s '%v'\n", introspection.GetTypeName(n), n)
//		rr = append(rr, node.Accept(v))
//	}
//	return rr
//}

// -----------------------------
// Override BaseExpressionsVisitor

func (v *Visitor) VisitTreeStart(ctx *TreeStartContext) interface{} {
	//print("START\n")
	//return v.VisitChildren(ctx)
	return ctx.Expr().Accept(v)
}

type Expr struct {
	Op    string
	Value int64
}

func (v *Visitor) VisitExpr(ctx *ExprContext) interface{} {
	//return v.VisitChildren(ctx)

	cnt, err := strconv.ParseInt(
		ctx.GetNum().GetText(), 10, 64,
	)
	if err != nil {
		panic(err)
	}
	op := ctx.GetOp().GetText()
	return Expr{
		op,
		cnt,
	}
}

func NewVisitor() Visitor {
	return Visitor{
		&BaseExpressionsVisitor{
			&antlr.BaseParseTreeVisitor{},
		},
	}
}
