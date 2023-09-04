package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xavier268/mylisp/mytest"
)

func TestLexer(t *testing.T) {

	tests := []string{
		// atoms, nbr, ' ,whitespaces
		"a",
		" a b",
		"2/3",
		"2/3  3/4",
		"  2/3  3/4  ",
		"'",
		" ' a",
		// stings
		`"a string" and
		   "another string"`,
		// comments
		`a b ; comment
		f
		; another comment
		d ;`,
		// mixing operators, numbers and identifiers
		")",
		"()",
		"( )",
		"//",
		"+=",
		"+ =",
		"==",
	}
	sb := new(strings.Builder)
	for i, test := range tests {
		fmt.Fprintln(sb)
		fmt.Fprintf(sb, "%d:\tinput : <%s>\n", i, test)

		lx := NewLexer(strings.NewReader(test), fmt.Sprintf("lexer test line %d", i))
		lval := new(mySymType)
		for k := 0; k < 30; k++ { // up to 30 tokens displayed

			ttype := lx.Lex(lval)
			if ttype == 0 {
				fmt.Fprintf(sb, "\tttype : 0\n")
				break
			}
			if ttype >= ' ' && ttype <= '~' {
				fmt.Fprintf(sb, "\tttype : %d (%c),\tlval : %#v\n", ttype, ttype, lval)
			} else {
				fmt.Fprintf(sb, "\tttype : %d (%X),\tlval : %#v\n", ttype, ttype, lval)
			}

		}

	}

	mytest.Verify(t, sb.String(), "lexer")
}
