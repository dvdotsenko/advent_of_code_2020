// Code generated from /mnt/src/Expressions.g4 by ANTLR 4.8. DO NOT EDIT.

package day8parser
import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 187, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 
	39, 9, 39, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 
	3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 94, 10, 5, 3, 5, 3, 5, 3, 6, 6, 6, 99, 10, 
	6, 13, 6, 14, 6, 100, 3, 6, 3, 6, 7, 6, 105, 10, 6, 12, 6, 14, 6, 108, 
	11, 6, 5, 6, 110, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 
	10, 3, 11, 3, 11, 7, 11, 122, 10, 11, 12, 11, 14, 11, 125, 11, 11, 3, 12, 
	6, 12, 128, 10, 12, 13, 12, 14, 12, 129, 3, 12, 3, 12, 3, 13, 3, 13, 3, 
	14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 
	3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 
	24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 
	3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 
	35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 2, 2, 
	40, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 
	23, 13, 25, 2, 27, 2, 29, 2, 31, 2, 33, 2, 35, 2, 37, 2, 39, 2, 41, 2, 
	43, 2, 45, 2, 47, 2, 49, 2, 51, 2, 53, 2, 55, 2, 57, 2, 59, 2, 61, 2, 63, 
	2, 65, 2, 67, 2, 69, 2, 71, 2, 73, 2, 75, 2, 77, 2, 3, 2, 32, 4, 2, 67, 
	92, 99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 
	3, 2, 50, 59, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2, 69, 69, 
	101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 103, 4, 2, 72, 72, 
	104, 104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74, 106, 106, 4, 2, 75, 75, 
	107, 107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77, 109, 109, 4, 2, 78, 78, 
	110, 110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80, 112, 112, 4, 2, 81, 81, 
	113, 113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83, 115, 115, 4, 2, 84, 84, 
	116, 116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86, 118, 118, 4, 2, 87, 87, 
	119, 119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89, 121, 121, 4, 2, 90, 90, 
	122, 122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92, 124, 124, 2, 166, 2, 3, 
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 
	3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 
	19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 3, 79, 3, 2, 2, 2, 
	5, 83, 3, 2, 2, 2, 7, 87, 3, 2, 2, 2, 9, 93, 3, 2, 2, 2, 11, 98, 3, 2, 
	2, 2, 13, 111, 3, 2, 2, 2, 15, 113, 3, 2, 2, 2, 17, 115, 3, 2, 2, 2, 19, 
	117, 3, 2, 2, 2, 21, 119, 3, 2, 2, 2, 23, 127, 3, 2, 2, 2, 25, 133, 3, 
	2, 2, 2, 27, 135, 3, 2, 2, 2, 29, 137, 3, 2, 2, 2, 31, 139, 3, 2, 2, 2, 
	33, 141, 3, 2, 2, 2, 35, 143, 3, 2, 2, 2, 37, 145, 3, 2, 2, 2, 39, 147, 
	3, 2, 2, 2, 41, 149, 3, 2, 2, 2, 43, 151, 3, 2, 2, 2, 45, 153, 3, 2, 2, 
	2, 47, 155, 3, 2, 2, 2, 49, 157, 3, 2, 2, 2, 51, 159, 3, 2, 2, 2, 53, 161, 
	3, 2, 2, 2, 55, 163, 3, 2, 2, 2, 57, 165, 3, 2, 2, 2, 59, 167, 3, 2, 2, 
	2, 61, 169, 3, 2, 2, 2, 63, 171, 3, 2, 2, 2, 65, 173, 3, 2, 2, 2, 67, 175, 
	3, 2, 2, 2, 69, 177, 3, 2, 2, 2, 71, 179, 3, 2, 2, 2, 73, 181, 3, 2, 2, 
	2, 75, 183, 3, 2, 2, 2, 77, 185, 3, 2, 2, 2, 79, 80, 5, 45, 23, 2, 80, 
	81, 5, 51, 26, 2, 81, 82, 5, 57, 29, 2, 82, 4, 3, 2, 2, 2, 83, 84, 5, 27, 
	14, 2, 84, 85, 5, 31, 16, 2, 85, 86, 5, 31, 16, 2, 86, 6, 3, 2, 2, 2, 87, 
	88, 5, 53, 27, 2, 88, 89, 5, 55, 28, 2, 89, 90, 5, 57, 29, 2, 90, 8, 3, 
	2, 2, 2, 91, 94, 5, 13, 7, 2, 92, 94, 5, 15, 8, 2, 93, 91, 3, 2, 2, 2, 
	93, 92, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 96, 5, 
	11, 6, 2, 96, 10, 3, 2, 2, 2, 97, 99, 5, 25, 13, 2, 98, 97, 3, 2, 2, 2, 
	99, 100, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 109, 
	3, 2, 2, 2, 102, 106, 5, 17, 9, 2, 103, 105, 5, 25, 13, 2, 104, 103, 3, 
	2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 
	2, 107, 110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 102, 3, 2, 2, 2, 109, 
	110, 3, 2, 2, 2, 110, 12, 3, 2, 2, 2, 111, 112, 7, 45, 2, 2, 112, 14, 3, 
	2, 2, 2, 113, 114, 7, 47, 2, 2, 114, 16, 3, 2, 2, 2, 115, 116, 7, 48, 2, 
	2, 116, 18, 3, 2, 2, 2, 117, 118, 7, 46, 2, 2, 118, 20, 3, 2, 2, 2, 119, 
	123, 9, 2, 2, 2, 120, 122, 9, 3, 2, 2, 121, 120, 3, 2, 2, 2, 122, 125, 
	3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 22, 3, 2, 
	2, 2, 125, 123, 3, 2, 2, 2, 126, 128, 9, 4, 2, 2, 127, 126, 3, 2, 2, 2, 
	128, 129, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 
	131, 3, 2, 2, 2, 131, 132, 8, 12, 2, 2, 132, 24, 3, 2, 2, 2, 133, 134, 
	9, 5, 2, 2, 134, 26, 3, 2, 2, 2, 135, 136, 9, 6, 2, 2, 136, 28, 3, 2, 2, 
	2, 137, 138, 9, 7, 2, 2, 138, 30, 3, 2, 2, 2, 139, 140, 9, 8, 2, 2, 140, 
	32, 3, 2, 2, 2, 141, 142, 9, 9, 2, 2, 142, 34, 3, 2, 2, 2, 143, 144, 9, 
	10, 2, 2, 144, 36, 3, 2, 2, 2, 145, 146, 9, 11, 2, 2, 146, 38, 3, 2, 2, 
	2, 147, 148, 9, 12, 2, 2, 148, 40, 3, 2, 2, 2, 149, 150, 9, 13, 2, 2, 150, 
	42, 3, 2, 2, 2, 151, 152, 9, 14, 2, 2, 152, 44, 3, 2, 2, 2, 153, 154, 9, 
	15, 2, 2, 154, 46, 3, 2, 2, 2, 155, 156, 9, 16, 2, 2, 156, 48, 3, 2, 2, 
	2, 157, 158, 9, 17, 2, 2, 158, 50, 3, 2, 2, 2, 159, 160, 9, 18, 2, 2, 160, 
	52, 3, 2, 2, 2, 161, 162, 9, 19, 2, 2, 162, 54, 3, 2, 2, 2, 163, 164, 9, 
	20, 2, 2, 164, 56, 3, 2, 2, 2, 165, 166, 9, 21, 2, 2, 166, 58, 3, 2, 2, 
	2, 167, 168, 9, 22, 2, 2, 168, 60, 3, 2, 2, 2, 169, 170, 9, 23, 2, 2, 170, 
	62, 3, 2, 2, 2, 171, 172, 9, 24, 2, 2, 172, 64, 3, 2, 2, 2, 173, 174, 9, 
	25, 2, 2, 174, 66, 3, 2, 2, 2, 175, 176, 9, 26, 2, 2, 176, 68, 3, 2, 2, 
	2, 177, 178, 9, 27, 2, 2, 178, 70, 3, 2, 2, 2, 179, 180, 9, 28, 2, 2, 180, 
	72, 3, 2, 2, 2, 181, 182, 9, 29, 2, 2, 182, 74, 3, 2, 2, 2, 183, 184, 9, 
	30, 2, 2, 184, 76, 3, 2, 2, 2, 185, 186, 9, 31, 2, 2, 186, 78, 3, 2, 2, 
	2, 9, 2, 93, 100, 106, 109, 123, 129, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "", "", "'+'", "'-'", "'.'", "','",
}

var lexerSymbolicNames = []string{
	"", "K_JMP", "K_ACC", "K_NOOP", "SIGNED_NUMERIC_LITERAL", "NUMERIC_LITERAL", 
	"PLUS", "MINUS", "DOT", "COMMA", "ATOM", "WHITESPACE",
}

var lexerRuleNames = []string{
	"K_JMP", "K_ACC", "K_NOOP", "SIGNED_NUMERIC_LITERAL", "NUMERIC_LITERAL", 
	"PLUS", "MINUS", "DOT", "COMMA", "ATOM", "WHITESPACE", "DIGIT", "A", "B", 
	"C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", 
	"R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

type ExpressionsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewExpressionsLexer(input antlr.CharStream) *ExpressionsLexer {

	l := new(ExpressionsLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Expressions.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExpressionsLexer tokens.
const (
	ExpressionsLexerK_JMP = 1
	ExpressionsLexerK_ACC = 2
	ExpressionsLexerK_NOOP = 3
	ExpressionsLexerSIGNED_NUMERIC_LITERAL = 4
	ExpressionsLexerNUMERIC_LITERAL = 5
	ExpressionsLexerPLUS = 6
	ExpressionsLexerMINUS = 7
	ExpressionsLexerDOT = 8
	ExpressionsLexerCOMMA = 9
	ExpressionsLexerATOM = 10
	ExpressionsLexerWHITESPACE = 11
)

