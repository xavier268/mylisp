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
		//"( un deux trois )",
		//"( un . deux )",
	}

	sb := new(strings.Builder)
	for i, tt := range tests {
		fmt.Fprintln(sb)
		fmt.Fprintf(sb, "%d: input : <%s>\n", i, tt)
		res, err := ParseString(tt, fmt.Sprintf("test string #%d", i))
		if err != nil {
			fmt.Fprintf(sb, "%d:\t ********** error: %v\n", i, err)
		}
		fmt.Fprintf(sb, "%d: result: <%s>\n", i, res)
	}

	mytest.Verify(t, sb.String(), "parser")
}
