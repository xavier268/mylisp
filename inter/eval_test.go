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

		"(+ 2 ( * 5 3 ) 2)",
		"(* 2 3 )",
		"(*)",

		// with quotes
		"'(+ 1 2 )",
		"( + '1 2 )",
		"( + 1 '2 )",
		"'a",
		"''a",
		"'''a",
		"(quote  )",
		"(quote 1 )",
		"(quote 1 2 )",
		"('+ 1 2 ) ",

		// let
		"(let ((x 2) (y 3))  (* x y)) ; --> 6 ",

		// // fail as not evaluable
		"(1) ; fail",
		"((+ 1 2 ) 3 ) ; fail",
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

func TestEvalDetail(t *testing.T) {

	tt := "(let ((x 2) (y 3))  (* x y))"

	sb := new(strings.Builder)
	fmt.Fprintln(sb)
	fmt.Fprintf(sb, "Input   : %s\n", tt)
	ttt, err := ParseString(tt, t.Name())
	if err != nil {
		fmt.Fprintf(sb, "******** %v\n", err)
	}
	fmt.Fprintf(sb, "Parsed  : %v\n", ttt)
	it := NewInter()
	res := it.Eval(ttt)
	fmt.Fprintf(sb, "Evalued : %v\n", ToString(res))

	mytest.Verify(t, sb.String(), "eval.detail")
}
