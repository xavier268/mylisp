package inter

import "fmt"

var ErrVariableUnassigned = fmt.Errorf("variable is bound but unassigned")
var ErrVariableNotBound = fmt.Errorf("variable is not bound")

// Evaluate a symbol as a Variable, returning its binding if any, possibly unassigned if variable exists but has no value yet (or nil value).
func evalVar(ev *Environnement, t Term) (res Term, ok bool) {
	if t == nil {
		return nil, false
	}
	if a, ok := t.(Symbol); ok {
		if tt, ok := ev.Get(a); ok {
			if tt == nil {
				return Error{ErrVariableUnassigned, a}, true
			}
			return tt, true
		}
	}
	return nil, false
}
