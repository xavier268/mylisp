package parser

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

type myLex struct {
	scanner    *bufio.Scanner
	ttype      int     // token type
	SourceName string  // source for the data, used to print relevant errors
	pos        int     // position reading (bytes)
	LastErr    []error // the errors encountered
	LastResult Term    // the last result from the parser
}

func NewLexer(input io.Reader, fileName string) *myLex {
	lx := &myLex{
		scanner:    bufio.NewScanner(input),
		ttype:      0,
		SourceName: fileName,
		pos:        0,
		LastErr:    []error{},
		LastResult: nil,
	}
	lx.scanner.Split(lx.splitFunc)
	return lx
}

var PAT_COMMENT = regexp.MustCompile("(?m)^;.*$")
var PAT_SPACE = regexp.MustCompile("^[ \t]+")
var PAT_STRING = regexp.MustCompile(`^"([^"]*)"`)
var PAT_NUMBER = regexp.MustCompile(`^(-)?(\d+)(/(\d+))?`)
var PAT_OPERATOR = regexp.MustCompile(`^[()'.]`)
var PAT_IDENT = regexp.MustCompile(`^[-+*/$_\pL\d]+`) // caution, this definition contains numbers !

// split function for the scanner
// type recognized is in lx.ttype
// responsible for eating comments and advancing line and pos
func (lx *myLex) splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	var pos []int

	if atEOF && len(data) == 0 {
		return 0, nil, bufio.ErrFinalToken
	}

	// eat white spaces
	pos = PAT_SPACE.FindIndex(data)
	if pos != nil && pos[0] == 0 { // found white space, eat  it
		lx.pos += pos[1]
		return pos[1], nil, nil
	}

	// eat comments
	pos = PAT_COMMENT.FindIndex(data)
	if pos != nil && pos[0] == 0 { // found comment
		if pos[1] <= len(data) || atEOF { // has remainer or end of file
			lx.pos += pos[1]
			return pos[1], nil, nil // eat comment
		} else {
			return 0, nil, nil // ask for more data
		}
	}

	// scan operators
	pos = PAT_OPERATOR.FindIndex(data)
	if pos != nil && pos[0] == 0 { // found operator
		if pos[1] <= len(data) || atEOF { // has remainer or end of file
			lx.ttype = int(data[0])
			lx.pos += pos[1]
			return pos[1], data[0:pos[1]], nil // return operator
		} else {
			return 0, nil, nil // ask for more data
		}
	}

	// scan strings
	pos = PAT_STRING.FindIndex(data)
	if pos != nil && pos[0] == 0 { // found string
		if pos[1] <= len(data) || atEOF { // has remainer or end of file
			lx.ttype = STRING
			lx.pos += pos[1]
			return pos[1], data[0:pos[1]], nil // return string
		} else {
			return 0, nil, nil // ask for more data
		}
	}

	// scan numbers
	pos = PAT_NUMBER.FindIndex(data)
	if pos != nil && pos[0] == 0 { // found number
		if pos[1] <= len(data) || atEOF { // has remainer or end of file
			lx.ttype = NUMBER
			lx.pos += pos[1]
			return pos[1], data[0:pos[1]], nil // return number
		} else {
			return 0, nil, nil // ask for more data
		}
	}

	// scan identifiers
	pos = PAT_IDENT.FindIndex(data)
	if pos != nil && pos[0] == 0 { // found identifier
		if pos[1] <= len(data) || atEOF { // has remainer or end of file
			lx.ttype = IDENT
			lx.pos += pos[1]
			return pos[1], data[0:pos[1]], nil // return identifier
		} else {
			return 0, nil, nil // ask for more data
		}
	}

	panic("not implemented")
}

// The parser calls this method to get each new token.
// The token type is the return value, and the token value is in the lval.
// The tokens types are : '(' ')' '.' '\” NUMBER IDENT STRING
// Comments starts with a ; until end of line, and are skipped.
// White spaces are skipped.
func (lx *myLex) Lex(lval *mySymType) int {
	if !lx.scanner.Scan() {
		if lx.scanner.Err() == nil { // eof
			return 0
		} else {
			lx.LastErr = append(lx.LastErr, fmt.Errorf("lexing error in  %s, pos %d  : %v", lx.SourceName, lx.pos, lx.scanner.Err()))
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
		lx.LastErr = append(lx.LastErr, fmt.Errorf("scan error in  %s, pos %d  :unknown scan type  %v for %s", lx.SourceName, lx.pos, lx.ttype, token))
		return ERROR
	}

}

// Required to satisfy interface with parser
func (lx *myLex) Error(s string) {
	lx.LastErr = append(lx.LastErr, fmt.Errorf("parse error in %s, pos %d : %v", lx.SourceName, lx.pos, s))
	fmt.Println(lx.LastErr)
}

// The myLexer interface is implemented by myLex, defined by generated parser.
var _ myLexer = new(myLex)
