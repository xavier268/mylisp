package inter

type Symbol struct {
	Value string
}

// IsProcedure implements Term.
func (Symbol) IsProcedure() bool {
	return false
}

// IsBool implements Term.
func (Symbol) IsBool() bool {
	return false
}

// IsNumber implements Term.
func (Symbol) IsNumber() bool {
	return false
}

// IsPair implements Term.
func (Symbol) IsPair() bool {
	return false
}

// IsString implements Term.
func (Symbol) IsString() bool {
	return false
}

// IsSymbol implements Term.
func (Symbol) IsSymbol() bool {
	return true
}

var _ Term = Symbol{}

func (t Symbol) String() string {
	return t.Value
}
