package inter

// Constructs a Pair object describing the error and its context.
// Format is : ( errorSymbol  errorMessageString context )
func TermError(err error, context Term) Term {
	return makeList(Symbol{Value: "error"}, Symbol{Value: err.Error()}, context)
}

// Make a single list with the provided terms.
// Returns nil, not Pair{}, if there are no parameters.
func makeList(t ...Term) Term {
	if t == nil {
		return nil
	}
	if len(t) == 1 {
		return Pair{
			Car: t[0],
			Cdr: nil,
		}
	}
	if len(t) >= 2 {
		return Pair{
			Car: t[0],
			Cdr: makeList(t[1:]...),
		}
	}
	panic(" case not implemented")
}

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
