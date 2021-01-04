// Code generated from /mnt/src/Expressions.g4 by ANTLR 4.8. DO NOT EDIT.

package day8parser // Expressions
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 13, 4, 
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 4, 
	2, 4, 2, 3, 3, 2, 3, 5, 2, 10, 2, 6, 3, 2, 2, 2, 4, 9, 3, 2, 2, 2, 6, 7, 
	5, 4, 3, 2, 7, 8, 7, 2, 2, 3, 8, 3, 3, 2, 2, 2, 9, 10, 9, 2, 2, 2, 10, 
	11, 7, 6, 2, 2, 11, 5, 3, 2, 2, 2, 2,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "'+'", "'-'", "'.'", "','",
}
var symbolicNames = []string{
	"", "K_JMP", "K_ACC", "K_NOOP", "SIGNED_NUMERIC_LITERAL", "NUMERIC_LITERAL", 
	"PLUS", "MINUS", "DOT", "COMMA", "ATOM", "WHITESPACE",
}

var ruleNames = []string{
	"treeStart", "expr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ExpressionsParser struct {
	*antlr.BaseParser
}

func NewExpressionsParser(input antlr.TokenStream) *ExpressionsParser {
	this := new(ExpressionsParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Expressions.g4"

	return this
}

// ExpressionsParser tokens.
const (
	ExpressionsParserEOF = antlr.TokenEOF
	ExpressionsParserK_JMP = 1
	ExpressionsParserK_ACC = 2
	ExpressionsParserK_NOOP = 3
	ExpressionsParserSIGNED_NUMERIC_LITERAL = 4
	ExpressionsParserNUMERIC_LITERAL = 5
	ExpressionsParserPLUS = 6
	ExpressionsParserMINUS = 7
	ExpressionsParserDOT = 8
	ExpressionsParserCOMMA = 9
	ExpressionsParserATOM = 10
	ExpressionsParserWHITESPACE = 11
)

// ExpressionsParser rules.
const (
	ExpressionsParserRULE_treeStart = 0
	ExpressionsParserRULE_expr = 1
)

// ITreeStartContext is an interface to support dynamic dispatch.
type ITreeStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTreeStartContext differentiates from other interfaces.
	IsTreeStartContext()
}

type TreeStartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTreeStartContext() *TreeStartContext {
	var p = new(TreeStartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionsParserRULE_treeStart
	return p
}

func (*TreeStartContext) IsTreeStartContext() {}

func NewTreeStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TreeStartContext {
	var p = new(TreeStartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionsParserRULE_treeStart

	return p
}

func (s *TreeStartContext) GetParser() antlr.Parser { return s.parser }

func (s *TreeStartContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TreeStartContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserEOF, 0)
}

func (s *TreeStartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TreeStartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TreeStartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionsVisitor:
		return t.VisitTreeStart(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExpressionsParser) TreeStart() (localctx ITreeStartContext) {
	localctx = NewTreeStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExpressionsParserRULE_treeStart)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.Expr()
	}
	{
		p.SetState(5)
		p.Match(ExpressionsParserEOF)
	}



	return localctx
}


// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token 

	// GetNum returns the num token.
	GetNum() antlr.Token 


	// SetOp sets the op token.
	SetOp(antlr.Token) 

	// SetNum sets the num token.
	SetNum(antlr.Token) 


	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op antlr.Token
	num antlr.Token
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionsParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionsParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) GetNum() antlr.Token { return s.num }


func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) SetNum(v antlr.Token) { s.num = v }


func (s *ExprContext) SIGNED_NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserSIGNED_NUMERIC_LITERAL, 0)
}

func (s *ExprContext) K_JMP() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserK_JMP, 0)
}

func (s *ExprContext) K_ACC() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserK_ACC, 0)
}

func (s *ExprContext) K_NOOP() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserK_NOOP, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionsVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExpressionsParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExpressionsParserRULE_expr)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(7)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ExprContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ExpressionsParserK_JMP) | (1 << ExpressionsParserK_ACC) | (1 << ExpressionsParserK_NOOP))) != 0)) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ExprContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(8)

		var _m = p.Match(ExpressionsParserSIGNED_NUMERIC_LITERAL)

		localctx.(*ExprContext).num = _m
	}



	return localctx
}


