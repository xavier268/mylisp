package inter

import "fmt"

var ErrEval = fmt.Errorf("cannot evaluate provided input")

// Eval a Term, recursively, returning the result as a Term.
func (it *Inter) Eval(t Term) Term {

	// Eval nil to nil
	if t == nil {
		return nil
	}

	// try to handle self evaluating forms that have no scope impact
	if tt, ok := DoSelfEval(t); ok {
		return tt
	}

	// set new local scope
	it.PushScope()
	defer it.PopScope()

	// do we have a defined variable ? If yes, return its binding without evaluating it.
	if tt, ok := it.DoVar(t); ok {
		return tt
	}

	// try to apply special predicates, if any.
	if tt, ok := it.DoSpecial(t); ok {
		return tt
	}

	// handle functions of the form  ( functor . args )
	if functor, ok := car(t).(Symbol); ok {
		if evfunc, ok := it.Get(functor); ok {
			// replace functor with its lamda expression definition, and evaluate the replacement.
			return it.Eval(Cell{
				Car: evfunc,
				Cdr: cdr(t),
			})

		}

	}

	// cannot evaluate, return error
	return TermError(ErrEval, t)

}

// DoSelfEval eval when we do not need access to the scopes.
func DoSelfEval(t Term) (res Term, ok bool) {
	switch a := t.(type) {
	case Number:
		return a, true
	case String:
		return a, true
	case Symbol:
		switch a.Value {
		case "true", "t":
			return t, true
		case "nil":
			return nil, true
		default:
			return nil, false
		}
	case Cell:
		switch {
		case a.Car == Symbol{Value: "error"}:
			return t, true
		case a.Car == Symbol{Value: "quote"}:
			return a.Cdr, true
		default:
			return nil, false
		}
	default:
		return nil, false
	}
}

// Evaluate a symbol that has a binding, returning its binding.
func (it *Inter) DoVar(t Term) (res Term, ok bool) {

	if a, ok := t.(Symbol); ok {
		if tt, ok := it.Get(a); ok {
			return tt, true
		}
	}
	return nil, false
}
