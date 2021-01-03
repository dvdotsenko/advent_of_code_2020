package day7parser

import (
	"aoc2020/pkg/introspection"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

type Expr struct {
	Operator string
	Left interface{}
	Right interface{}
}
func (e Expr) String() string {
	return fmt.Sprintf("(%s %v %s)", e.Left, e.Operator, e.Right)
}

type ValueType int
const (
	INT_TYPE ValueType = iota + 1
	FLOAT_TYPE
	STR_TYPE
)

type Literal struct {
	Value int64 // interface{}
	LiteralValue string
	Type ValueType
}
func (l Literal) String() string {
	return l.LiteralValue
}


type R struct {
	Value interface{}
	Error error
}

type Visitor struct {
	// This embed is redundant and is used only to express "inheritance" tree
	// All methods from this and its parent must be reimplemented here at top level
	// because almost all of them are NOOP by default
	*BaseExpressionsVisitor
}

// -----------------------------
// Override BaseParseTreeVisitor

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	fmt.Printf("Visit : %s '%v'\n", introspection.GetTypeName(tree), tree.(antlr.ParseTree).GetText())
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
	print("START\n")
	//return v.VisitChildren(ctx)
	return ctx.GetValue().Accept(v)
}

func (v *Visitor) VisitExpr(ctx *ExprContext) interface{} {
	fmt.Printf("Expression :'%v'\n", ctx.GetText())
	//return v.VisitChildren(ctx)

	if v := ctx.NUMBER(); v != nil {
		t := v.GetText()
		vv, e := strconv.ParseInt(t, 10, 64)
		if e != nil {
			panic(fmt.Sprintf("Value should be Int, but found '%s'", t))
		}
		return Literal{
			vv,
			t,
			INT_TYPE,
		}
	}

	// if not number, must be experssion
	l := ctx.GetLeft()
	r := ctx.GetRight()
	o := ctx.GetOp()

	return Expr{
		o.GetText(),
		l.Accept(v),
		r.Accept(v),
	}
}


func NewVisitor() Visitor {
	return Visitor{
		&BaseExpressionsVisitor{
			&antlr.BaseParseTreeVisitor{},
		},
	}
}
