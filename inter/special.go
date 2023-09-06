package inter

// Handlling of special functors

// Function to transform a Term that starts or contains a special predicate.
// The full Term is provided as input, eg ( lambda ( X Y ) ( + X Y ))
// If ok is false, res is irrelevant, most likely nil, t should be kept, and further eval proceed.
// If ok is true, res is the transformed Term, and eval is finished.
// Errors are a special case of transformation, with ok = true.
// The functor that triggered the call is not necessary checked again to be the car of t, but it should.
type TransformFunc func(Term) (res Term, ok bool)

// Execute special builtin predicates, if any, and return true.
// Return false if nothing to do, and eval should continue.
// In false, then the result is set to nil, to avoid unecessary copy.
// Special should include quote, lambda, error, ...
// Parameter evaluation is never done implicitely, and must be managed by the special implementation, if needed.
func (it *Inter) DoSpecial(t Term) (result Term, ok bool) {

	if t == nil {
		return nil, true
	}
	switch t := t.(type) {
	// No specials here
	case String:
		return nil, false
	case Number:
		return nil, false
	case Symbol: // input is a
		// Reserved future use - maybe atom predicates later , like 'quit' ?
		return nil, false
	case Cell: // input is ( a . b )
		if t.Car == nil { // ( . b )
			return nil, false
		}

		panic("to be continued")
		// TODO // (  + - * /  ...)
		// TODO // ( ( lamda ...))

	}
	panic("to be continued")
}
