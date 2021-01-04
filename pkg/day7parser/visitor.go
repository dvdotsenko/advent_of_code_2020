package day7parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
	"strings"
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

type Component struct {
	Count int64
	Label string
}

type Expr struct {
	Label      string
	Components []Component
}

func (v *Visitor) VisitExpr(ctx *ExprContext) interface{} {
	//return v.VisitChildren(ctx)

	l := ctx.Label().Accept(v).(string)
	cc := []Component{}
	for _, c := range ctx.AllComponent() {
		cv := c.Accept(v)
		if cv != nil {
			cc = append(cc, cv.(Component))
		}
	}
	return Expr{
		l,
		cc,
	}
}

func (v *Visitor) VisitComponent(ctx *ComponentContext) interface{} {
	//return v.VisitChildren(ctx)

	if ctx.K_NO() != nil {
		return nil
	}

	cnt, err := strconv.ParseInt(ctx.GetCount().GetText(), 10, 64)
	if err != nil {
		panic(err)
	}

	l := ctx.Label().Accept(v).(string)

	return Component{
		cnt,
		l,
	}
}

func (v *Visitor) VisitLabel(ctx *LabelContext) interface{} {
	//return v.VisitChildren(ctx)
	ss := []string{}
	for _, s := range ctx.AllATOM() {
		ss = append(ss, s.GetText())
	}
	return strings.Join(ss, " ")
}

func NewVisitor() Visitor {
	return Visitor{
		&BaseExpressionsVisitor{
			&antlr.BaseParseTreeVisitor{},
		},
	}
}
