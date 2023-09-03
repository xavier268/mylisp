package parser

import (
	"bufio"
	"fmt"
	"io"
)

type myLex struct {
	scanner    *bufio.Scanner
	ttype      int     // token type
	SourceName string  // source for the data, used to print relevant errors
	line, col  int     // position reading
	LastErr    []error // the errors encountered
	LastResult Term    // the last result from the parser
}

func NewLexer(input io.Reader, fileName string) *myLex {
	lx := &myLex{
		scanner:    bufio.NewScanner(input),
		ttype:      0,
		SourceName: fileName,
		line:       0,
		col:        0,
		LastErr:    []error{},
		LastResult: nil,
	}
	lx.scanner.Split(lx.splitFunc)
	return lx
}

// split function for the scanner
// type recognized is in lx.ttype
// responsible for eating comments and advancing line and pos
func (lx *myLex) splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// set token type in a field in myLex ? Using regexp ?
	panic("not implemented")
}

// The parser calls this method to get each new token.
// The token type is the return value, and the token value is in the lval.
// The tokens types are : '(' ')' '.' '\” NUMBER IDENT STRING
// Comments starts with a ; until end of line, and are skipped.
func (lx *myLex) Lex(lval *mySymType) int {
	if !lx.scanner.Scan() {
		if lx.scanner.Err() == nil { // eof
			return 0
		} else {
			lx.LastErr = append(lx.LastErr, fmt.Errorf("lexing error in  %s, line %d - col %d : %v", lx.SourceName, lx.line, lx.col, lx.scanner.Err()))
			return ERROR
		}
	}
	token := lx.scanner.Text()
	switch lx.ttype {
	case 0:
		lval.value = nil
		return 0
	case '(', ')', '.', '\'':
		lval.value = nil
		return lx.ttype
	case IDENT:
		lval.value = Atom{
			Value: token,
		}
		return IDENT
	case NUMBER:
		lval.value = Number{ // TODO : convert token to rational
			Num: 0,
			Den: 0,
		}
		return NUMBER
	case STRING:
		lval.value = String{
			Value: token,
		}
		return STRING
	default:
		lx.LastErr = append(lx.LastErr, fmt.Errorf("scan error in  %s, line %d - col %d :unknown scan type  %v for %s", lx.SourceName, lx.line, lx.col, lx.ttype, token))
		return ERROR
	}

}

// Required to satisfy interface with parser
func (lx *myLex) Error(s string) {
	lx.LastErr = append(lx.LastErr, fmt.Errorf("parse error in %s, line %d - col %d : %v", lx.SourceName, lx.line, lx.col, s))
	fmt.Println(lx.LastErr)
}

// The myLexer interface is implemented by myLex, defined by generated parser.
var _ myLexer = new(myLex)
