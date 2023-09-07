package inter

// gives car if it makes sense, or nil.
// never fail, never error.
func car(t Term) Term {
	if t == nil {
		return nil
	}
	if c, ok := t.(Pair); ok {
		return c.Car
	}
	return nil
}

func cdr(t Term) Term {
	if t == nil {
		return nil
	}
	if c, ok := t.(Pair); ok {
		return c.Cdr
	}
	return nil
}

// keep the compiler happy ..
var _ = [...]any{car, cdr, caar, cadr, cdar, cddr}

func caar(t Term) Term {
	return car(car(t))
}

func cadr(t Term) Term {
	return car(cdr(t))
}

func cdar(t Term) Term {
	return cdr(car(t))
}

func cddr(t Term) Term {
	return cdr(cdr(t))
}
