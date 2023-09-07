package parser

type Term interface {
	String() string
	IsPair() bool
	IsSymbol() bool
	IsString() bool
	IsNumber() bool
	IsBool() bool
	IsProcedure() bool
}

var _ Term = Pair{}
var _ Term = Symbol{}
var _ Term = String{}
var _ Term = Number{}
var _ Term = Bool{}

// a convenience string function, thaht handles nil Term in the correct way.
func ToString(t Term) string {
	if t == nil {
		return "<nil>"
	}
	return t.String()
}
