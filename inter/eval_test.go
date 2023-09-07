package inter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xavier268/mylisp/mytest"
)

func TestEval(t *testing.T) {

	tests := []string{

		"",
		"1",
		"(+)",
		"(+ 1 )",
		"( + 1 2)",
		"( + 1/2 3/4)",
		"( + 1 (+ 2/3 1/3 ))",
		"(+ (+ 1 2) )",

		// with quotes
		"'(+ 1 2 )",
		"( + '1 2 )",
		"( + 1 '2 )",

		// "( eval (+ 1 2)	)",

		"('+ 1 2 ) ; fail as not evaluable",
		"( eval '+) ; fail",
		"( eval ('+ )); should be ok",
		"( (eval '+) 1 2 ) ; should work",

		// // fail as not evaluable
		"(1)",

		"((+ 1 2 ) 3 )",
	}

	sb := new(strings.Builder)

	for i, tt := range tests {
		fmt.Fprintln(sb)
		fmt.Fprintf(sb, "%3d: Input   : %s\n", i, tt)
		ttt, err := ParseString(tt, fmt.Sprintf("%s - test # %d", t.Name(), i))
		if err != nil {
			fmt.Fprintf(sb, "******** %v\n", err)
		}
		fmt.Fprintf(sb, "%3d: Parsed  : %v\n", i, ttt)
		it := NewInter()
		res := it.Eval(ttt)
		fmt.Fprintf(sb, "%3d: Evalued : %v\n", i, res)
		// fmt.Fprintln(sb, it)
	}

	mytest.Verify(t, sb.String(), "eval")

}
