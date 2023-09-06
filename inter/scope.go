package inter

import "fmt"

type Scope struct {
	bindings map[Symbol]Term
	parent   *Scope
}

func (it *Inter) PushScope() {
	it.local = &Scope{
		bindings: map[Symbol]Term{},
		parent:   it.local,
	}
	if it.global == nil {
		it.global = it.local // for initialization
	}
}

var ErrCannotPopScope = fmt.Errorf("cannot pop scope, scope stack is empty")

func (it *Inter) PopScope() error {
	if it.local == nil || it.local.parent == nil {
		return ErrCannotPopScope
	} else {
		it.local = it.local.parent
		return nil
	}
}

// get symbol binding in all loacl or global scopes, nil if not found.
func (it *Inter) Get(sym Symbol) (res Term, ok bool) {
	if it.local == nil {
		panic("scope is nil")
	}
	for s := it.local; s != nil; s = s.parent {
		if t, ok := s.bindings[sym]; ok {
			return t, true
		}
	}
	return nil, false // not found
}

// Set symbol in local scope
func (it *Inter) Set(sym Symbol, t Term) {
	if it.local == nil {
		panic("scope is nil")
	}
	it.local.bindings[sym] = t
}

// set symbol in global scope
func (it *Inter) SetGlobal(sym Symbol, t Term) {
	if it.global == nil {
		panic("scope is nil")
	}
	it.global.bindings[sym] = t
}

// binds a tree of arguments to a target tree.
// only the symbols matching Terms in the same position in the target tree are bound.
func (it *Inter) Bind(args, target Term) {

	if args == nil {
		return
	}

	switch a := args.(type) {
	case Number: // ignore
		return
	case String: // ignore
		return
	case Symbol:
		it.Bind(a, target)
		return
	case Cell:
		if tc, ok := target.(Cell); ok {
			it.Bind(a.Car, tc.Car)
			it.Bind(a.Cdr, tc.Cdr)
			return
		} else {
			return // no  match
		}
	default:
		panic("case not implemented")
	}
}
