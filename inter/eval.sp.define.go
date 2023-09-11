package inter

// Syntax : (define variable expression) or ( define variable )
// This implementation will define the variable in the environement where define is executed.
// TODO - so maybe, we don't need the notion of global/local environement ?
// Expression IS EVALUATED before binding.

func init() {
	// (define variable expression)
	// expression WILL BE EVALUATED before binding, even if flag is false, but not variable.
	Register("define", false, spDefine)
	// (set! variable expression)
	// expression WILL BE EVALUATED before binding, even if flag is false, but not variable.
	Register("set!", false, spSetBang)
}

// (define variable expression) or ( define variable )
// Always return nil
// Expression IS EVALUATED before being stored.
// If variable is not bound, define binds variable to a new location in the CURRENT environment
// before performing the assignment.
// If variable is already bound, define may replace or shadow replaces the existing binding, depending on environement.
func spDefine(it *Inter, t Term) Term { // if ( define var value ) was called, t is now ( var value)
	if car(t) != nil && car(t).IsSymbol() {
		it.current.Set(car(t).(Symbol), it.Eval(cadr(t)))
	}
	return nil
}

// (set! var value) or (set! var)
// Return nil or an Error Term.
// Will look for an EXISTING variable called var. If none found, error.
// Expression IS EVALUATED before being stored.
// It is an error to perform a set! on an unbound variable.
// If you omit expression, the variable becomes unassigned; an attempt to reference such a variable is an error.
func spSetBang(it *Inter, t Term) Term { // t is now ( var value)
	if car(t) != nil && car(t).IsSymbol() {
		err := it.current.SetBang(car(t).(Symbol), it.Eval(cadr(t)))
		if err != nil {
			return Error{err, t}
		}
	}
	return nil
}
