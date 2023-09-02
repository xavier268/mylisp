package parser

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type myLex struct {
	reader     bufio.Reader
	SourceName string
	line, col  int     // position reading
	token      string  // accumulator for token string
	LastErr    []error // the errors encountered
	LastResult Term    // the last result from the parser
}

func NewLexer(input io.Reader, fileName string) *myLex {
	return &myLex{
		reader:     bufio.NewReader(input), // buffer input to read rune by rune.
		SourceName: fileName,
		line:       0,
		col:        0,
		token:      "",
		LastErr:    []error{},
		LastResult: nil,
	}
}

// rune classifier.
// Alpha are returned as 'A',
// space as ' '
// digits as '0'
// "';().\n are returned as is,
// eof is 0, unknown is -1.
func rtype(r rune) int {
	switch {
	case strings.ContainsRune(";()'.\n\"/+-*", r):
		return int(r)
	case unicode.IsSpace(r): // \n already handled above
		return ' '
	case unicode.IsDigit(r):
		return '0'
	case r == 0:
		return 0
	case unicode.IsLetter(r) || strings.ContainsRune("_$+-*", r):
		return 'A'
	default:
		return ERROR
	}
}

// The parser calls this method to get each new token.
// The token type is the return value, and the token value is in the lval.
// The tokens types are : '(' ')' '.' '\'' NUMBER IDENT STRING
// Comments starts with a ; until end of line, and are skipped.

func (lx *myLex) Lex(lval *mySymType) int {

	lval.value = nil // reset value to nil
	var (
		r    rune
		err  error
		size int
	)

start:

	r = lx.NextRune()
	switch rtype(r) {

	case '(', ')':
		return int(r)

		// TODO - continuer !!
	}

	goto start

	panic("unimplemented")
}

func (lx *myLex) NextRune() rune {
	r, _, err := lx.reader.ReadRune()
	lx.col += 1
	if r == '\n' {
		lx.line += 1
		lx.col = 0
	}
	if err != nil {
		lx.LastErr = append(lx.LastErr, err)
		return ERROR
	}
	return r
}

// Required to satisfy interface.
func (lx *myLex) Error(s string) {
	lx.LastErr = append(lx.LastErr, fmt.Errorf("error in %s, line %d - col %d : %v", lx.SourceName, lx.line, lx.col, s))
	fmt.Println(lx.LastErr)
}

// The myLexer interface is implemented by myLex, defined by generated parser.
var _ myLexer = new(myLex)
