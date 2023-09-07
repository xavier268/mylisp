package inter

import "fmt"

var ErrEval = fmt.Errorf("cannot evaluate provided input")

// Eval a Term, recursively, returning the result as a Term.
// An expression to be evaluated may be a literal, a variable reference, a special form, or a procedure call.
// TODO - review how environments are managed ?
func (it *Inter) Eval(t Term) Term {

	if it.quit {
		return nil
	}

	fmt.Println("DEBUG : eval de : ", t)

	// If litteral, done.
	// nil and Pair{} are litterals.
	if IsLitteral(t) {
		return t
	}

	// if a bound variable, evaluate it in the current environement and return its value.
	// Error if variable is bound but unassigned.
	if tt, ok := EvalVar(it.current, t); ok {
		return tt
	}

	// if special form, try to apply special predicate
	if car(t) != nil && car(t).IsSymbol() && IsKeyword(car(t).(Symbol).Value) { // we have ( symbol ... )
		return it.EvalSpecial(t)
	}

	// if the car is not nil, evaluate and replace the car (and only the car) to see if we can continue.
	// That will cover the case of quoted functors, as well as translatinf functors into their lambada expression.
	// Do not evaluate yet the other parameters.
	if ca := car(t); ca != nil && ca.IsSymbol() {
		if tt, ok := EvalVar(it.current, ca); ok {
			return it.Eval(Pair{
				Car: tt,
				Cdr: cdr(t), // args not evaluated yet
			})
		}
	}

	// cannot evaluate, return error
	return Error{ErrEval, t}

}
