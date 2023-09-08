package inter

import (
	"fmt"
	"slices"
	"strings"
)

type Environnement struct {
	bindings map[Symbol]Term
	parent   *Environnement
}

func (s Environnement) String() string {
	sb := new(strings.Builder)

	// collect keys
	keys := make([]Symbol, 0, len(s.bindings))
	for s2 := range s.bindings {
		keys = append(keys, s2)
	}
	// sort keys
	slices.SortFunc(keys, func(a, b Symbol) int { return strings.Compare(a.Value, b.Value) })
	// display scope
	for _, k := range keys {
		v := s.bindings[k]
		if v == nil {
			fmt.Fprintf(sb, "\t%s = nil\n", k)
		} else {
			fmt.Fprintf(sb, "\t%s = %s\n", k, v)
		}
	}
	return sb.String()
}

func (ev *Environnement) Fork() *Environnement {
	return &Environnement{
		bindings: map[Symbol]Term{},
		parent:   ev,
	}
}

func (it *Inter) PushEnv() {
	it.current = &Environnement{
		bindings: map[Symbol]Term{},
		parent:   it.current,
	}
	if it.global == nil {
		it.global = it.current // for initialization
	}
}

var ErrCannotPopENv = fmt.Errorf("cannot pop environement,  stack is empty")

func (it *Inter) PopEnv() error {
	if it.current == nil || it.current.parent == nil {
		return ErrCannotPopENv
	} else {
		it.current = it.current.parent
		return nil
	}
}

// get symbol binding in environements, recursively. Ok if found.
func (ev *Environnement) Get(sym Symbol) (res Term, ok bool) {
	if ev == nil {
		panic("scope is nil")
	}
	for s := ev; s != nil; s = s.parent {
		if t, ok := s.bindings[sym]; ok {
			return t, true
		}
	}
	return nil, false // not found
}

// Set symbol in current environement, creating it if needed
func (ev *Environnement) Set(sym Symbol, t Term) {
	if ev == nil {
		panic("scope is nil")
	}
	if (sym == Symbol{}) {
		panic("symbol has no name")
	}
	ev.bindings[sym] = t
}

// Try to find the specified variable in one of the environements,
// and set this specific variable in this specific environement to the value.
// Variable is NEVER created if not found, but an error will be returned.
// A existing variable will NEVER be shadowed, but will be changed.
func (ev *Environnement) SetBang(sym Symbol, t Term) error {
	if ev == nil {
		return ErrVariableNotBound
	}
	if _, ok := ev.bindings[sym]; ok {
		ev.bindings[sym] = t
		return nil
	}
	pv := ev.parent
	if pv == nil {
		return ErrVariableNotBound
	} else {
		return pv.SetBang(sym, t)
	}
}

// binds a tree of arguments to a target tree.
// only the symbols matching Terms in the same position in the target tree are bound.
func (ev *Environnement) Bind(args, target Term) {

	if args == nil {
		return
	}

	switch a := args.(type) {
	case Number: // ignore
		return
	case String: // ignore
		return
	case Symbol:
		ev.Set(a, target)
		return
	case Pair:
		if tc, ok := target.(Pair); ok {
			ev.Bind(a.Car, tc.Car)
			ev.Bind(a.Cdr, tc.Cdr)
			return
		} else {
			return // no  match
		}
	default:
		panic("case not implemented")
	}
}
