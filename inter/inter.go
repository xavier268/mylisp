package inter

import (
	"fmt"

	"github.com/xavier268/mylisp/parser"
)

type Inter struct {
}

// type aliases
type Term = parser.Term
type String = parser.String
type Number = parser.Number
type Atom = parser.Atom
type Cell = parser.Cell

// Pseudo error to request successful end of evaluation.
var ErrEndOfEval = fmt.Errorf("eval is done")

// Variable is not currently bound.
var ErrNotBound = fmt.Errorf("variable not bound")

// Eval a Term, returning the result as a Term.
func (it *Inter) Eval(t Term) Term {
	var ok, ch bool
	var err error

startEval:

	for {
		if t == nil {
			return nil
		}

		// String evaluates to itself
		if _, ok = t.(String); ok {
			return t
		}
		// Number evaluates to itself
		if _, ok = t.(Number); ok {
			return t
		}

		// apply special predicates, if any, before any substitution had a chance to happen.
		tt, ch, err := it.DoSpecial(t)
		if err == ErrEndOfEval { // abort further evaluation with success
			return tt
		}
		if err != nil { // stop & report error
			return ReportError(err, t)
		}
		if ch {
			t = tt
			continue startEval
		}

		// Ok, here no changed by builtin special, let's continue....

		// Do we have a single Atom ?
		if a, ok := t.(Atom); ok {
			// is it bound to something ?
			if tt, ok := it.GetVariable(a); ok {
				// bound ! Replace and loop evaluation
				return tt // Don't restart eval, what if X is bound to Y and Y is bound to X ? Endless loop !
			}
			// not bound, we're done !
			return a
		}

		// Do we have a pair/list ?
		
		panic("to do")

		}

		// loop
	}

}

// Execute special builtin predicates, if any.
// Evaluation will continue, unless error is returned.
// The special ErrEndOfEval serves to request immediate but successful termination of the evaluation.
// Special should include tick, lambda, error, ...
func (it *Inter) DoSpecial(t Term) (Term, bool, error) {
	panic("not implemented	")
}

// Reports error as an error with context, as an object that will print and read easily
func ReportError(err error, context Term) Term {
	// returns : ( ( error "the error text") context )
	return Cell{
		Car: Cell{
			Car: Atom{Value: "error"},
			Cdr: String{Value: err.Error()},
		},
		Cdr: context,
	}
}

// Returns the Term currently bound to the variable v, telling if found.
func (it *Inter) GetVariable(v Atom) (t Term, found bool) {
	panic("not implemented")
}
