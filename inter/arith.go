package inter

import (
	"fmt"

	"github.com/xavier268/mylisp/parser"
)

var ErrNotANumber = fmt.Errorf("not a number")

// Compute sum of all numbers in Term, which must be in a list as in ( 1  X  (+ 3 4)  4), whose memebres evaluate as Numbers.
// The '+' operateur is not part of the submited list.
// Return an error term if one of them is not a number.
// Remember that specials are in charge of evaluating their input parameters.
func (it *Inter) Plus(t Term) Term {

	if (t == nil || t == Cell{}) { // nil or empty list
		return parser.NumberZero
	}

	if n, ok := it.Eval(car(t)).(Number); ok { // eval the first argument, make sure result is a number
		if m, ok := it.Plus(cdr(t)).(Number); ok { // eval the rest, ensure result is a number
			return n.Plus(m) // add the numbers
		}
	}
	return TermError(ErrNotANumber, t)
}
