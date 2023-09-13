package inter

import "fmt"

// Should only be called with a term of the form ( keyword ...)
func (it *Inter) evalSpecial(t Term) Term {

	kw := car(t).(Symbol).Value
	kwd := keywordsDefinitions[kw]

	args := cdr(t)

	if kwd.evalArgFirst {
		args = it.evalListArgs(args)
	}
	if kwd.primitive != nil {
		return kwd.primitive(it, args)
	}
	return Error{ErrEval, t}
}

var ErrInvalidArgumentListFormat = fmt.Errorf("invalid argument list format")

// eval the list of arguments
func (it *Inter) evalListArgs(t Term) Term {

	if t == nil {
		return nil
	}

	if !t.IsPair() {
		return Error{ErrInvalidArgumentListFormat, t}
	}

	return Pair{
		Car: it.Eval(car(t)),
		Cdr: it.evalListArgs(cdr(t)),
	}
}
