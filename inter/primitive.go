package inter

// Handlling of special functors

// Execute special builtin predicates, if any, and return true.
// Return false if nothing to do, and eval should continue.
// In false, then the result is set to nil, to avoid unecessary copy.
// Special should include quote, lambda, error, ...
// Parameter evaluation is never done implicitely, and must be managed by the special implementation, if needed.
func (it *Inter) DoSpecial(t Term) (result Term, ok bool) {

	if t == nil {
		return nil, true
	}
	switch { // all specials are selected here
	case car(t) == Symbol{Value: "quote"}:
		return cdr(t), true
	case car(t) == Symbol{Value: "eval"}:
		return it.Eval(cdr(t)), true
	case car(t) == Symbol{Value: "+"}:
		return it.Plus(cdr(t)), true
	}
	// default  ..
	return nil, false // no special found, no applicable
}
