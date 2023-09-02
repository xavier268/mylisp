package parser

import (
	"fmt"
	"io"
	"text/scanner"
)

type myLex struct {
	s          scanner.Scanner
	LastErr    []error // the errors encountered
	LastResult Term    // the last result from the parser
}

func NewLexer(input io.Reader, fileName string) *myLex {
	lx := new(myLex)
	lx.s.Init(input)
	lx.s.Filename = fileName
	lx.s.Mode = scanner.ScanIdents | scanner.ScanInts |  scanner.ScanStrings 
	return lx
}

// The parser calls this method to get each new token.
// The token type is the return value, and the token value is in the lval.
func (lx *myLex) Lex(lval *mySymType) int {

	lval.value = nil // reset value to nil
	tk := lx.s.Scan()

	switch tk {
	case scanner.EOF:
		return 0 // special eof type for parser
	case '(',')','.','\':
		return int(tk)

	}

	panic("unimplemented")
}

// Required to satisfy interface.
func (lx *myLex) Error(s string) {
	lx.LastErr = append(lx.LastErr, fmt.Errorf("error in %s, line %d : %v", lx.s.Filename, lx.s.Line, s))
	fmt.Println(lx.LastErr)
}

// The myLexer interface is implemented by myLex, defined by generated parser.
var _ myLexer = new(myLex)
