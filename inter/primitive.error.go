package inter

// Constructs a Pair object describing the error and its context.
// Format is : ( errorSymbol  errorMessageString context )
func TermError(err error, context Term) Term {
	return makeList(Symbol{Value: "error"}, Symbol{Value: err.Error()}, context)
}

// Term is an error if it is of the form ( errorSymbol ... )
func IsError(t Term) bool {
	return t != nil && car(t).IsSymbol() && car(t).(Symbol).Value == "error"
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
