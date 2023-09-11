package inter

import "fmt"

var ErrEval = fmt.Errorf("cannot evaluate provided input")

// Eval a Term, recursively, returning the result as a Term.
// An expression to be evaluated may be a literal, a variable reference, a special form, or a procedure call.
// Only the primitives or procedures that require it should fork the environement.
func (it *Inter) Eval(t Term) Term {

	if DEBUG >= 2 {
		fmt.Println("DEBUG : evaluating : ", ToString(t))
	}

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

	// if special form with Var or Keyword, try to apply special predicate to it.
	if car(t) != nil && car(t).IsSymbol() { // t is ( symbol ... )
		if tt, ok := it.current.Get(car(t).(Symbol)); ok && tt != nil { // symbol is a var, get its value  !
			return it.Eval(Pair{ // evaluate resulting form
				Car: tt,
				Cdr: cdr(t),
			})
		}
		if IsKeyword(car(t).(Symbol).Value) { // t is  ( keyword ... )
			return it.EvalSpecial(t)
		}
	}

	// if the functor is a procedure, call it.
	if car(t) != nil && car(t).IsProcedure() { // t is ( proc arg1 arg2 ... )
		return it.EvalProcedure(t)
	}

	// if the functor could be a form like ( (something ... )  ... )
	if car(t) != nil && car(t).IsPair() && caar(t).IsSymbol() { // we have something like ( ( funct ... )  ... )
		return it.Eval(
			Pair{
				Car: it.Eval(car(t)),
				Cdr: cdr(t),
			})
	}

	// cannot evaluate, return error
	return Error{ErrEval, t}

}
