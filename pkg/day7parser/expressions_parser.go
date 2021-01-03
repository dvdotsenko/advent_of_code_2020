// Code generated from /mnt/src/Expressions.g4 by ANTLR 4.8. DO NOT EDIT.

package day7parser // Expressions
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 40, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 3, 5, 
	3, 15, 10, 3, 3, 3, 3, 3, 3, 3, 6, 3, 20, 10, 3, 13, 3, 14, 3, 21, 3, 4, 
	3, 4, 3, 4, 3, 4, 5, 4, 28, 10, 4, 5, 4, 30, 10, 4, 3, 4, 3, 4, 3, 4, 3, 
	5, 6, 5, 36, 10, 5, 13, 5, 14, 5, 37, 3, 5, 2, 2, 6, 2, 4, 6, 8, 2, 4, 
	3, 2, 5, 6, 3, 2, 9, 10, 2, 40, 2, 10, 3, 2, 2, 2, 4, 14, 3, 2, 2, 2, 6, 
	29, 3, 2, 2, 2, 8, 35, 3, 2, 2, 2, 10, 11, 5, 4, 3, 2, 11, 12, 7, 2, 2, 
	3, 12, 3, 3, 2, 2, 2, 13, 15, 5, 8, 5, 2, 14, 13, 3, 2, 2, 2, 14, 15, 3, 
	2, 2, 2, 15, 16, 3, 2, 2, 2, 16, 17, 7, 6, 2, 2, 17, 19, 7, 7, 2, 2, 18, 
	20, 5, 6, 4, 2, 19, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 19, 3, 2, 2, 
	2, 21, 22, 3, 2, 2, 2, 22, 5, 3, 2, 2, 2, 23, 24, 7, 3, 2, 2, 24, 30, 7, 
	4, 2, 2, 25, 27, 7, 8, 2, 2, 26, 28, 5, 8, 5, 2, 27, 26, 3, 2, 2, 2, 27, 
	28, 3, 2, 2, 2, 28, 30, 3, 2, 2, 2, 29, 23, 3, 2, 2, 2, 29, 25, 3, 2, 2, 
	2, 30, 31, 3, 2, 2, 2, 31, 32, 9, 2, 2, 2, 32, 33, 9, 3, 2, 2, 33, 7, 3, 
	2, 2, 2, 34, 36, 7, 11, 2, 2, 35, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 
	35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 9, 3, 2, 2, 2, 7, 14, 21, 27, 29, 
	37,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "'.'", "','",
}
var symbolicNames = []string{
	"", "K_NO", "K_OTHER", "K_BAG", "K_BAGS", "K_CONTAIN", "NUMERIC_LITERAL", 
	"DOT", "COMMA", "ATOM", "WHITESPACE",
}

var ruleNames = []string{
	"treeStart", "expr", "component", "label",
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
	ExpressionsParserK_NO = 1
	ExpressionsParserK_OTHER = 2
	ExpressionsParserK_BAG = 3
	ExpressionsParserK_BAGS = 4
	ExpressionsParserK_CONTAIN = 5
	ExpressionsParserNUMERIC_LITERAL = 6
	ExpressionsParserDOT = 7
	ExpressionsParserCOMMA = 8
	ExpressionsParserATOM = 9
	ExpressionsParserWHITESPACE = 10
)

// ExpressionsParser rules.
const (
	ExpressionsParserRULE_treeStart = 0
	ExpressionsParserRULE_expr = 1
	ExpressionsParserRULE_component = 2
	ExpressionsParserRULE_label = 3
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
		p.SetState(8)
		p.Expr()
	}
	{
		p.SetState(9)
		p.Match(ExpressionsParserEOF)
	}



	return localctx
}


// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *ExprContext) K_BAGS() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserK_BAGS, 0)
}

func (s *ExprContext) K_CONTAIN() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserK_CONTAIN, 0)
}

func (s *ExprContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *ExprContext) AllComponent() []IComponentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComponentContext)(nil)).Elem())
	var tst = make([]IComponentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComponentContext)
		}
	}

	return tst
}

func (s *ExprContext) Component(i int) IComponentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComponentContext)
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
	p.SetState(12)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ExpressionsParserATOM {
		{
			p.SetState(11)
			p.Label()
		}

	}
	{
		p.SetState(14)
		p.Match(ExpressionsParserK_BAGS)
	}
	{
		p.SetState(15)
		p.Match(ExpressionsParserK_CONTAIN)
	}
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == ExpressionsParserK_NO || _la == ExpressionsParserNUMERIC_LITERAL {
		{
			p.SetState(16)
			p.Component()
		}


		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IComponentContext is an interface to support dynamic dispatch.
type IComponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCount returns the count token.
	GetCount() antlr.Token 


	// SetCount sets the count token.
	SetCount(antlr.Token) 


	// IsComponentContext differentiates from other interfaces.
	IsComponentContext()
}

type ComponentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	count antlr.Token
}

func NewEmptyComponentContext() *ComponentContext {
	var p = new(ComponentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionsParserRULE_component
	return p
}

func (*ComponentContext) IsComponentContext() {}

func NewComponentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentContext {
	var p = new(ComponentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionsParserRULE_component

	return p
}

func (s *ComponentContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentContext) GetCount() antlr.Token { return s.count }


func (s *ComponentContext) SetCount(v antlr.Token) { s.count = v }


func (s *ComponentContext) K_BAGS() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserK_BAGS, 0)
}

func (s *ComponentContext) K_BAG() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserK_BAG, 0)
}

func (s *ComponentContext) DOT() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserDOT, 0)
}

func (s *ComponentContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserCOMMA, 0)
}

func (s *ComponentContext) K_NO() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserK_NO, 0)
}

func (s *ComponentContext) K_OTHER() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserK_OTHER, 0)
}

func (s *ComponentContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(ExpressionsParserNUMERIC_LITERAL, 0)
}

func (s *ComponentContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *ComponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ComponentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionsVisitor:
		return t.VisitComponent(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExpressionsParser) Component() (localctx IComponentContext) {
	localctx = NewComponentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ExpressionsParserRULE_component)
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExpressionsParserK_NO:
		{
			p.SetState(21)
			p.Match(ExpressionsParserK_NO)
		}
		{
			p.SetState(22)
			p.Match(ExpressionsParserK_OTHER)
		}


	case ExpressionsParserNUMERIC_LITERAL:
		{
			p.SetState(23)

			var _m = p.Match(ExpressionsParserNUMERIC_LITERAL)

			localctx.(*ComponentContext).count = _m
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == ExpressionsParserATOM {
			{
				p.SetState(24)
				p.Label()
			}

		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(29)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ExpressionsParserK_BAG || _la == ExpressionsParserK_BAGS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(30)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ExpressionsParserDOT || _la == ExpressionsParserCOMMA) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionsParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionsParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) AllATOM() []antlr.TerminalNode {
	return s.GetTokens(ExpressionsParserATOM)
}

func (s *LabelContext) ATOM(i int) antlr.TerminalNode {
	return s.GetToken(ExpressionsParserATOM, i)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionsVisitor:
		return t.VisitLabel(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExpressionsParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ExpressionsParserRULE_label)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == ExpressionsParserATOM {
		{
			p.SetState(32)
			p.Match(ExpressionsParserATOM)
		}


		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


