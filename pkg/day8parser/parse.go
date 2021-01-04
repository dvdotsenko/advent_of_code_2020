package day8parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
}

func (ce *ErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	//fmt.Fprintf(os.Stderr, "ERROR: line %d:%d %s\n", line, column, msg)
	panic(fmt.Sprintf("ERROR: line %d:%d %s\n", line, column, msg))
}

func Parse(s string) interface{} {
	io := antlr.NewInputStream(s)

	lex := NewExpressionsLexer(io)
	lex.RemoveErrorListeners()
	lex.AddErrorListener(new(ErrorListener))
	//lex.AddErrorListener(antlr.NewDiagnosticErrorListener(false))

	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)

	p := NewExpressionsParser(tokens)
	p.RemoveErrorListeners()
	p.AddErrorListener(new(ErrorListener))
	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))

	visitor := NewVisitor()
	tree := p.TreeStart()
	return visitor.Visit(tree)
}
