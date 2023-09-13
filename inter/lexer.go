package inter

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
	LastResult []Term  // the  result from the parser
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

var myPAT_COMMENT = regexp.MustCompile("(?m)^;.*$")
var myPAT_SPACE = regexp.MustCompile(`^[\s]+`)
var myPAT_STRING = regexp.MustCompile(`^"([^"]*)"`)
var myPAT_NUMBER = regexp.MustCompile(`^(-)?(\d+)(/(\d+))?`)    // should be tested before IDENT
var myPAT_BOOL = regexp.MustCompile(`^(#f|#t)`)                 // should be tested before IDENT
var myPAT_IDENT = regexp.MustCompile(`^[-+*/=$_\pL\d@?!%&<>]+`) // IDENT match this AND are neither number nor operator

var PAT_OPERATOR = regexp.MustCompile(`^[()'.]$`) // tell if an ident is an operator, once identified as IDENT

// split function for the scanner
// type recognized is in lx.ttype
// responsible for eating comments and advancing line and pos
func (lx *myLex) splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	var pos []int

	if atEOF && len(data) == 0 {
		lx.ttype = 0
		return 0, nil, nil
	}

	// eat white spaces
	pos = myPAT_SPACE.FindIndex(data)
	switch {
	case pos == nil || pos[0] != 0:
		// no number, ignore
	case pos[1] == len(data) && atEOF: // extend until end of file
		lx.ttype = SKIP
		lx.pos += pos[1]
		return pos[1], data[0:pos[1]], bufio.ErrFinalToken // final token
	case pos[1] == len(data) && !atEOF: // could be missing something
		return 0, nil, nil // ask for more data
	case pos[1] < len(data): // found, and there will be more
		lx.ttype = SKIP
		lx.pos += pos[1]
		return pos[1], data[0:pos[1]], nil // return
	default:
		panic("case not implemented")
	}

	// eat comments
	pos = myPAT_COMMENT.FindIndex(data)
	switch {
	case pos == nil || pos[0] != 0:
		// no number, ignore
	case pos[1] == len(data) && atEOF: // extend until end of file
		lx.ttype = SKIP
		lx.pos += pos[1]
		return pos[1], data[0:pos[1]], bufio.ErrFinalToken // final token
	case pos[1] == len(data) && !atEOF: // could be missing something
		return 0, nil, nil // ask for more data
	case pos[1] < len(data): // found, and there will be more
		lx.ttype = SKIP
		lx.pos += pos[1]
		return pos[1], data[0:pos[1]], nil // return
	default:
		panic("case not implemented")
	}

	// scan strings
	pos = myPAT_STRING.FindIndex(data)
	switch {
	case pos == nil || pos[0] != 0:
		// no number, ignore
	case pos[1] == len(data) && atEOF: // extend until end of file
		lx.ttype = STRING
		lx.pos += pos[1]
		return pos[1], data[0:pos[1]], bufio.ErrFinalToken // final token
	case pos[1] == len(data) && !atEOF: // could be missing something
		return 0, nil, nil // ask for more data
	case pos[1] < len(data): // found, and there will be more
		lx.ttype = STRING
		lx.pos += pos[1]
		return pos[1], data[0:pos[1]], nil // return number
	default:
		panic("case not implemented")
	}

	// immediately process parenthesis, quote and period
	if data[0] == '(' || data[0] == ')' || data[0] == '.' || data[0] == '\'' {
		lx.pos += 1
		lx.ttype = int(data[0])
		return 1, data[0:1], nil
	}

	// scan numbers
	pos = myPAT_NUMBER.FindIndex(data)
	switch {
	case pos == nil || pos[0] != 0:
		// no number, ignore
	case pos[1] == len(data) && atEOF: // extend until end of file
		lx.ttype = NUMBER
		lx.pos += pos[1]
		return pos[1], data[0:pos[1]], bufio.ErrFinalToken // final token
	case pos[1] == len(data) && !atEOF: // could be missing something
		return 0, nil, nil // ask for more data
	case pos[1] < len(data): // found, and there will be more
		lx.ttype = NUMBER
		lx.pos += pos[1]
		return pos[1], data[0:pos[1]], nil // return number
	default:
		panic("case not implemented")
	}

	// scan bools
	pos = myPAT_BOOL.FindIndex(data)
	switch {
	case pos == nil || pos[0] != 0:
		// no bool, ignore
	case pos[1] == len(data) && atEOF: // extend until end of file
		lx.ttype = BOOL
		lx.pos += pos[1]
		return pos[1], data[0:pos[1]], bufio.ErrFinalToken // final token
	case pos[1] == len(data) && !atEOF: // could be missing something ( not really for bool, but okay ...)
		return 0, nil, nil // ask for more data
	case pos[1] < len(data): // found, and there will be more
		lx.ttype = BOOL
		lx.pos += pos[1]
		return pos[1], data[0:pos[1]], nil // return bool
	default:
		panic("case not implemented")
	}

	// scan identifiers
	pos = myPAT_IDENT.FindIndex(data)
	switch {
	case pos == nil || pos[0] != 0:
		// no number, ignore
	case pos[1] == len(data) && atEOF: // extend until end of file
		lx.ttype = IDENT
		lx.pos += pos[1]
		return pos[1], data[0:pos[1]], bufio.ErrFinalToken // final token
	case pos[1] == len(data) && !atEOF: // could be missing something
		return 0, nil, nil // ask for more data
	case pos[1] < len(data): // found, and there will be more
		lx.ttype = IDENT
		lx.pos += pos[1]
		return pos[1], data[0:pos[1]], nil // return number
	default:
		panic("case not implemented")
	}

	return 0, nil, fmt.Errorf("token not recognized by lexer : %s", data)
}

// The parser calls this method to get each new token.
// The token type is the return value, and the token value is in the lval.
// The tokens types are : '(' ')' '.' '\” NUMBER IDENT STRING
// Comments starts with a ; until end of line, and are skipped.
// White spaces are skipped.
func (lx *myLex) Lex(lval *mySymType) int {

	for {
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
		case SKIP:
			continue // ignore and continue
		case IDENT:
			lval.value = Symbol{
				Value: token,
			}
			if PAT_OPERATOR.MatchString(token) {
				op := string(token)
				d := int(op[0])
				// fmt.Printf("DEBUG : type as char : %d  %c\n", d, d)
				return d // need to capture the full, multibyte, first character
			}
			return IDENT
		case '(', ')', '\'', '.':
			return int(token[0])
		case BOOL: // boolean
			if token == "#t" {
				lval.value = Bool{true}
				return BOOL
			} else {
				lval.value = Bool{false}
				return BOOL
			}
		case NUMBER:
			n, err := NumberFromString(token)
			if err != nil {
				lx.LastErr = append(lx.LastErr, fmt.Errorf("scan error in  %s, pos %d  :unknown number for %s :  %v ", lx.SourceName, lx.pos, token, err))
				return ERROR
			}
			lval.value = n
			return NUMBER
		case STRING:
			lval.value = String{
				Value: token[1 : len(token)-1], // strip quotation marks
			}
			return STRING
		default:
			lx.LastErr = append(lx.LastErr, fmt.Errorf("scan error in  %s, pos %d  :unknown scan type  %v for %s", lx.SourceName, lx.pos, lx.ttype, token))
			return ERROR
		}
	}
}

// Required to satisfy interface with parser
func (lx *myLex) Error(s string) {
	lx.LastErr = append(lx.LastErr, fmt.Errorf("parse error in %s, pos %d : %v", lx.SourceName, lx.pos, s))
	fmt.Println(lx.LastErr)
}

// The myLexer interface is implemented by myLex, defined by generated parser.
var _ myLexer = new(myLex)
