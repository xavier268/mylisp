package inter

import "fmt"

var ErrEval = fmt.Errorf("cannot evaluate provided input")

// Eval a Term, recursively, returning the result as a Term.
// An expression to be evaluated may be a literal, a variable reference, a special form, or a procedure call.
// TODO - review how environments are managed ?
func (it *Inter) Eval(t Term) Term {

	fmt.Println("DEBUG : eval de : ", t)

	// Eval nil to nil
	if t == nil {
		return nil
	}

	// If litteral, done.
	if tt, ok := EvalLiteral(t); ok {
		return tt
	}

	// if a bound variable, evaluate it in the current environement and return its value.
	// Error if variable is bound but unassigned.
	if tt, ok := it.EvalVar(t); ok {
		return tt
	}

	// set new local scope
	it.PushEnv()
	defer it.PopEnv()

	// try to apply special predicates, if any, before evaluating
	if tt, ok := it.DoSpecial(t); ok {
		return tt
	}

	// if the car is not nil, evaluate and replace, to see if we can continue.
	// That will cover the case of quoted functors, as well as translatinf functors into their lambada expression.
	// Do not evaluate yet the other parameters.
	if ca := car(t); ca != nil {
		if tt, ok := it.EvalVar(ca); ok {
			return it.Eval(Pair{
				Car: tt,
				Cdr: nil,
			})
		}
	}

	// cannot evaluate, return error
	return TermError(ErrEval, t)

}

// EvalLiteral eval when we do not need access to the scopes.
// Includes literal and error forms.
func EvalLiteral(t Term) (res Term, ok bool) {
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
	case Symbol: // error litteral or other keywords
		switch a.Value {
		case "error":
			return t, true
		}
	case Pair: // error form litteral
		if a.Car != nil && (a.Car == Symbol{Value: "error"}) {
			return t, true
		}
	}
	return nil, false
}

var ErrVariableUnassigned = fmt.Errorf("variable is bound but unassigned")

// Evaluate a symbol as a Variable, returning its binding if any, possibly unassigned if variable exists but has no value yet (or nil value).
func (it *Inter) EvalVar(t Term) (res Term, ok bool) {

	if a, ok := t.(Symbol); ok {
		if tt, ok := it.Get(a); ok {
			if tt == nil {
				return TermError(ErrVariableUnassigned, a), true
			}
			return tt, true
		}
	}
	return nil, false
}
