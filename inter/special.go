package inter

import (
	"fmt"

	"github.com/xavier268/mylisp/parser"
)

// Handlling of special functors

// Function to transform a Term that starts or contains a special predicate.
// The full Term is provided as input, eg ( lambda ( X Y ) ( + X Y ))
// If ok is false, res is irrelevant, t should be kept, and further eval proceed.
// If ok is true, res is the transformed Term, and eval is finished.
// Errors are a special case of transformation, with ok = true.
// The functor that triggered the call is not necessary checked again to be the car of t, but it should.
type TransformFunc func(Term) (res Term, ok bool)

// Execute special builtin predicates, if any, and return true.
// Return false if nothing to do, and eval should continue.
// In false, then the result is set to nil, to avoid unecessary copy.
// Special should include tick, lambda, error, ...
// Parameter evaluation is never done, if not explicitely requested.
func (it *Inter) DoSpecial(t Term) (result Term, ok bool) {

	if t == nil {
		return nil, true
	}

	switch t := t.(type) {
	// No specials here
	case String:
		return nil, false
	case Number:
		return nil, false
	case Atom:
		// Reserved future use - maybe atom predicates later ?
		return nil, false
	case Cell:
		if t.Car == nil {
			return nil, false
		}
		if a, ok := t.Car.(Atom); ok {
			// Special predicates could apply here ?
			sp, ok := it.specials[a]
			if !ok {
				// no predicate
				return nil, false
			} else {
				// apply corresponding predicate
				return sp(a)
			}
		}
		return nil, false
	default:
		panic("internal error - case should not happen")
	}
}

// Initialize the special builtin predicates
func (it *Inter) InitSpecials() {
	it.specials = map[parser.Atom]TransformFunc{
		Atom{Value: "tick"}:   it.specTick,
		Atom{Value: "lambda"}: it.specLambda,
		Atom{Value: "+"}:      it.specPlus,
		Atom{Value: "-"}:      it.specMinus,
		Atom{Value: "*"}:      it.specMul,
		Atom{Value: "/"}:      it.specDiv,
	}
}

// return the tick'd Term without evaluation, or an error.
func (it *Inter) specTick(t Term) (res Term, ok bool) {
	if c, ok := t.(Cell); ok {
		return c.Cdr, true // stop eval !
	} else { // not a cell, return error
		err := fmt.Errorf("tick: argument is not a cell")
		return TermError(err, t), true // stop eval
	}
}

func (it *Inter) specLambda(t Term) (res Term, ok bool) {
	panic("not implementd")
}

func (it *Inter) specPlus(t Term) (res Term, ok bool) {
	panic("not implementd")
}

func (it *Inter) specMinus(t Term) (res Term, ok bool) {
	panic("not implementd")
}

func (it *Inter) specMul(t Term) (res Term, ok bool) {
	panic("not implementd")
}

func (it *Inter) specDiv(t Term) (res Term, ok bool) {
	panic("not implementd")
}
