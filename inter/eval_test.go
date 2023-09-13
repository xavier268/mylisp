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

		"(*)",
		"(* 3 )",
		"(* 2 3 )",
		"(* 2 3 4)",

		"(/)",
		"( / 2 )",
		" ( / 2 3 ) ",
		" ( / 2 3 4  )",
		" ( / 2 3 4 5 ) ",

		"(-)",
		"(- 2 )",
		"(- 2 3 )",
		"(- 2 3 4  )",
		"(- 2 3 4 5 ) ",

		// list of symbols
		"( a) ; should fail ",
		" ( a b ) ; should fail",

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

		// let & define
		"(let ((x 2) (y 3))  (* x y)) ; --> 6 ",
		"(let ((x 2) )  (+ x y)) ; fail, y is unbound ",
		"(let ())",
		"(let () ( define a ( + 1 2 )) a ) ; --> 3",
		"( let ( ( x 3 ) ( y 0 ) ) ( let ( ( x 444 ) ) ( set! y x ) ) ( + y x ) ) ;--> 447, because shadowing x",

		// // fail as not evaluable
		"(1) ; fail",
		"((+ 1 2 ) 3 ) ; fail",

		// car, cons, cdr, list
		"(list)",
		"(list 'a )",
		"(list a ) ; fail because a unbound",
		"(car( a b)) ; fail to evaluate (a b)",
		"(car ' (a b)) ",
		"(cdr ' (a b)) ",
		"(car ' ()) ",
		"(cdr ' ()) ",
		"(cons)",

		// let shadowing
		`(let  (
			( x 3 ) ( y 0 ))
			 ( let (( x 444 )) ; shadowing x
					 ( set! y x) 
					 ) ; end of inner let
						   ( + y x) )`,

		// lambda
		"(lambda ( x y ) ( + x y ))",
		"(lambda ( x y ) ( + x y ) ( * x y ))",
		"(lambda ( x ) ( + x y ) )",
		"(lambda () ( + y) )",
		"(lambda () ) ; do nothing ",

		// lambda , applied
		" (( lambda ( x y ) ( + x x y) ) 2 3 ) ; 7",

		// begin
		"(begin 1 ( + 1 2 ) ( * 3 3 ))",

		// dump-env
		"( dump- env) ",
		"( let ( ( x 3 )) ( dump-env))",

		// keywords list
		"( keywords)",
	}

	sb := new(strings.Builder)

	for i, tt := range tests {
		fmt.Fprintln(sb)
		fmt.Fprintf(sb, "%3d: Input   : %s\n", i, tt)
		ttt, err := Parse1String(tt, fmt.Sprintf("%s - test # %d", t.Name(), i))
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

	DEBUG = 5

	tt := "(begin 1 ( + 1 2 ) ( * 3 3 ))"
	sb := new(strings.Builder)
	fmt.Fprintln(sb)
	fmt.Fprintf(sb, "Input   : %s\n", tt)
	ttt, err := Parse1String(tt, t.Name())
	if err != nil {
		fmt.Fprintf(sb, "******** %v\n", err)
	}
	fmt.Fprintf(sb, "Parsed  : %v\n", ttt)
	it := NewInter()
	res := it.Eval(ttt)
	fmt.Fprintf(sb, "Evalued : %v\n", ToString(res))

	mytest.Verify(t, sb.String(), "eval.detail")
}
