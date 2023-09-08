package inter

// Binds a variable to a value in the current environement.
// Syntax : (define variable expression) or ( define variable )
// Expression IS evaluated before binding.
// A top-level definition, (define variable expression), has essentially the same effect as this assignment
// expression, if variable is bound:(set! variable expression)
// If variable is not bound, however, define binds variable to a new location in the current environment
// before performing the assignment (it is an error to perform a set! on an unbound variable).
// If you omit expression, the variable becomes unassigned; an attempt to reference such a variable is an error.
func init() {
	Register("define", false, spDefine) // (define variable expression)
	Register("set!", false, spSetBang)  // (set! variable expression)
}

// Always return nil
// Expression IS EVALUATED before being stored.
func spDefine(it *Inter, t Term) Term { // if ( define var value ) was called, t is now ( var value)
	if car(t) != nil && car(t).IsSymbol() {
		it.current.Set(car(t).(Symbol), it.Eval(cadr(t)))
	}
	return nil
}

// Return nil or an Error Term.
// Expression IS EVALUATED before being stored.
func spSetBang(it *Inter, t Term) Term {
	if car(t) != nil && car(t).IsSymbol() {
		err := it.current.SetBang(car(t).(Symbol), it.Eval(cadr(t)))
		if err != nil {
			return Error{err, t}
		}
	}
	return nil
}
