package inter

import "fmt"

var ErrEval = fmt.Errorf("cannot evaluate provided input")

// Eval a Term, recursively, returning the result as a Term.
// TODO - review how environments are managed ?
func (it *Inter) Eval(t Term) Term {

	fmt.Println("DEBUG : eval de : ", t)

	// Eval nil to nil
	if t == nil {
		return nil
	}

	// try to handle self evaluating forms that evaluate to themselves.
	if tt, ok := DoSelfEval(t); ok {
		return tt
	}

	// set new local scope
	it.PushEnv()
	defer it.PopEnv()

	// do we have a defined variable ? If yes, return its binding without evaluating it.
	if tt, ok := it.DoVar(t); ok {
		return tt
	}

	// try to apply special predicates, if any, before evaluating
	if tt, ok := it.DoSpecial(t); ok {
		return tt
	}

	// if the car is not nil, evaluate and replace, to see if we can continue.
	// That will cover the case of quoted functors, as well as translatinf functors into their lambada expression.
	// Do not evaluate yet the other parameters.
	if ca := car(t); ca != nil {
		if tt, ok := it.DoVar(ca); ok {
			return it.Eval(Pair{
				Car: tt,
				Cdr: nil,
			})
		}
	}

	// cannot evaluate, return error
	return TermError(ErrEval, t)

}

// DoSelfEval eval when we do not need access to the scopes.
func DoSelfEval(t Term) (res Term, ok bool) {
	if t == nil {
		return t, true
	}
	switch a := t.(type) {
	case Bool:
		return a, true
	case Number:
		return a, true
	case String:
		return a, true
	case Symbol: // none, for now ?
	case Pair: // none for the moment ?
	}
	return nil, false
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
