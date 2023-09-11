package inter

import "fmt"

func (it *Inter) EvalProcedure(t Term) Term { // t contains ( proc arg1 arg2 ...) , args list could be empty.

	p := car(t).(Procedure)

	// TODO : do we need to eval the args for the procedure first, if yes, in which environnement ?

	// change environnement to forked closure.
	oldenv := it.current      // save old environement
	it.current = p.Env.Fork() // fork the closure environement, make it current
	defer func() {            // restore old environement on exit
		it.current = oldenv
	}()

	// bind arguments to formals
	err := bindArgs(it.current, p.Formals, cdr(t))
	if err != nil {
		return Error{err, t}
	}

	// evaluate the body/bodies in the closure environement
	var result Term
	for bodies := p.Body; bodies != nil && (bodies != Pair{}); bodies = cdr(bodies) {

		exp := car(bodies)
		result = it.Eval(exp)
	}

	return result

}

// binds formals arguments to actual arguments
// return error if mismatch, or nil if ok
func bindArgs(env *Environnement, formals Term, args Term) error {
	// formals contains ( var1 var2  ...), args contains ( arg1 arg2 ...),
	if formals == nil || (formals == Pair{}) {
		// done
		return nil
	}
	if formals.IsSymbol() {
		env.Set(formals.(Symbol), args) // binds to the list of args
		return nil
	}
	if formals.IsPair() && car(formals).IsSymbol() {
		env.Set(car(formals).(Symbol), car(args)) // binds to the first arg
		return bindArgs(env, cdr(formals), cdr(args))
	}
	return fmt.Errorf("procedure : arguments do not match formal parameters")

}
