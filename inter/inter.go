package inter

type Inter struct {
	specials map[Atom]TransformFunc
}

func NewInter() *Inter {
	it := new(Inter)
	it.InitSpecials()
	return it
}

func (it *Inter) GetVariable(a Atom) (resTerm Term, ok bool) {
	panic("not implemented	")
}

// Constructs a Cell object describing the error and its context.
// Format is : ( errorTagAtom  errorMessageString context )
func TermError(err error, context Term) Term {
	return MakeList(Atom{Value: "error"}, Atom{Value: err.Error()}, context)
}

// Make a single list with the provided terms.
// Returns nil, not Cell{}, if there are no parameters.
func MakeList(t ...Term) Term {
	if t == nil {
		return Cell{Car: nil, Cdr: nil}
	}
	if len(t) == 1 {
		return Cell{
			Car: t[0],
			Cdr: nil,
		}
	}
	if len(t) >= 2 {
		return Cell{
			Car: t[0],
			Cdr: MakeList(t[1:]...),
		}
	}
	panic(" case not implemented")
}
