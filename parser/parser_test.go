package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xavier268/mylisp/mytest"
)

func TestParser(t *testing.T) {

	tests := []string{
		"a",
		"2/3",
		`"a string"`,
		" ' 2/3  ",
		" a b ; should fail beacause multiple atoms not allowed outside list",
		"' ; should fail because ' must be followed by a Term",
	}

	sb := new(strings.Builder)
	for i, tt := range tests {
		fmt.Fprintln(sb)
		fmt.Fprintf(sb, "%d: input : <%s>\n", i, tt)
		res, err := ParseString(tt, fmt.Sprintf("test string #%d", i))
		if err != nil {
			fmt.Fprintf(sb, "%d:\t ********** error: %v\n", i, err)
		}
		fmt.Fprintf(sb, "%d: result raw: <%#v>\n", i, res)
		fmt.Fprintf(sb, "%d: result str: <%s>\n", i, ToString(res))
	}

	mytest.Verify(t, sb.String(), "parser")
}

func TestParserList(t *testing.T) {

	tests := []string{
		// simple lists
		"() ; should fail, because () is not ( )",
		"( )",
		"'( )",
		"( un deux )",
		"( un . deux )",
		"( un ' deux )",
		"'( un  deux )",
		"( un deux trois )",

		// nested lists
		"( un ( deux ) trois )",
		"( un ( deux () ) trois )",

		// using pairs to construct lists
		" (un . ( deux . ( trois . () ) ) )",
		" (un . ( deux . ( trois .  ) ) )",
	}

	sb := new(strings.Builder)
	for i, tt := range tests {
		fmt.Fprintln(sb)
		fmt.Fprintf(sb, "%d: input : <%s>\n", i, tt)
		res, err := ParseString(tt, fmt.Sprintf("test string #%d", i))
		if err != nil {
			fmt.Fprintf(sb, "%d:\t ********** error: %v\n", i, err)
		}
		fmt.Fprintf(sb, "%d: result raw: <%#v>\n", i, res)
		fmt.Fprintf(sb, "%d: result str: <%s>\n", i, ToString(res))
	}

	mytest.Verify(t, sb.String(), "parser.list")
}
