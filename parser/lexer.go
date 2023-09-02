package parser

import (
	"fmt"
	"text/scanner"
)

type myLex struct {
	s          scanner.Scanner
	LastErr    []error
	LastResult Term
}

// The parser calls this method to get each new token.
// The token type is the return value, and the token value is in the lval.
func (*myLex) Lex(lval *mySymType) int {
	panic("unimplemented")
}

// Required to satisfy interface.
func (lx *myLex) Error(s string) {
	lx.LastErr = append(lx.LastErr, fmt.Errorf("error in %s, line %d : %v", lx.s.Filename, lx.s.Line, s))
	fmt.Println(lx.LastErr)
}

// The myLexer interface is implemented by myLex, defined by generated parser.
var _ myLexer = new(myLex)
