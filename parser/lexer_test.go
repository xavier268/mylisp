package parser

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/xavier268/mylisp/mytest"
)

func TestLexer(t *testing.T) {

	tests := []string{
		"a",
		" a b",
		"2/3",
		"2/3  3/4",
		"  2/3  3/4  ",
		"'",
		" ' a",
	}
	sb := new(strings.Builder)
	for i, test := range tests {
		fmt.Fprintln(sb)
		fmt.Fprintf(sb, "%d:\tinput : <%s>\n", i, test)

		lx := NewLexer(strings.NewReader(test), fmt.Sprintf("lexer test line %d", i))
		lval := new(mySymType)
		for k := 0; k < 5; k++ {

			ttype := lx.Lex(lval)
			if unicode.IsLetter(rune(ttype)) {
				fmt.Fprintf(sb, "\tttype : %d (%c),\tlval : %#v\n", ttype, rune(ttype), lval)
			} else {
				fmt.Fprintf(sb, "\tttype : %d (%X),\tlval : %#v\n", ttype, ttype, lval)
			}

		}

	}

	mytest.Verify(t, sb.String(), "lexer")
}
