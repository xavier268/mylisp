package inter

// Eval a Term, recursively, returning the result as a Term.
func (it *Inter) Eval(t Term) Term {

	// try to apply special predicates, if any, before any substitution has a chance to happen.
	if tt, ok := it.DoSpecial(t); ok {
		return tt
	}

	// Do we have a single Atom ?
	if a, ok := t.(Atom); ok {
		// is it bound to something ?
		if tt, ok := it.GetVariable(a); ok {
			// bound ! Replace and return bind value.
			return tt
		}
		// not bound, we're done !
		return a
	}

	// Do we have a pair/list ?
	if c, ok := t.(Cell); ok {
		if c.Car == nil {
			return t
		}
		// eval functor and argument
		return Cell{Car: it.Eval(c.Car), Cdr: it.Eval(c.Cdr)} // if f was a function, it has been replaced now by a lambda expression.
	}

	// nothing else to evaluate, return the original term
	return t

}
