package inter

type Term interface {
	String() string
	IsPair() bool
	IsSymbol() bool
	IsString() bool
	IsNumber() bool
	IsBool() bool
	IsProcedure() bool
	IsError() bool
}

var _ Term = Pair{}
var _ Term = Symbol{}
var _ Term = String{}
var _ Term = Number{}
var _ Term = Bool{}
var _ Term = Procedure{}
var _ Term = Error{}
var _ Term = nil

// a convenience string function, that handles nil Term.
func ToString(t Term) string {
	if t == nil {
		return "<nil>"
	}
	return t.String()
}
