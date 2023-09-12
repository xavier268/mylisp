package inter

func init() {
	Register("list", true, spList, []string{
		" ( list arg1 arg2 ...) creates a lit : ( arg1 arg2 ...).",
	})
	Register("car", true, spCar, []string{
		"car returns the first element of the pair, or nil if not found",
	})
	Register("cdr", true, spCdr, []string{
		"cdr returns the rest of the pair, or nil if not found",
	})
	Register("cons", true, spCons, []string{
		"( cons obj1 ob2 )",
		"cons returns a new pair, (obj1.obj2)",
	})
}

// ( list arg1 arg2 ...) creates a lit : ( arg1 arg2 ...).
// arguments are evaluated first.
func spList(_ *Inter, t Term) Term { // t will contain (arg1 arg2 ...)
	if t == nil {
		return Pair{}
	}
	if !t.IsPair() {
		return Error{ErrInvalidArgumentListFormat, t}
	}
	return t
}

// ( car pair ) -> return the first element of the pair, or nil if not found
func spCar(_ *Inter, t Term) Term { // t contains (pair)
	return caar(t)
}

// ( cdr pair ) -> return the rest of the pair, or nil if not found
func spCdr(_ *Inter, t Term) Term { // t contains (pair)
	return cdar(t)
}

// ( cons obj1 obj2) -> return a new pair, (obj1.obj2)
// There can be less than 2 object, of more. If more, the rest is ignored (but evaluated).
func spCons(_ *Inter, t Term) Term { // t contains (obj1 obj2)
	return Pair{
		Car: car(t),
		Cdr: cadr(t),
	}
}
