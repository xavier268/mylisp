package parser

type Bool struct {
	Value bool
}

var _ Term = Bool{}

func (b Bool) String() string {
	if b.Value {
		return "#t"
	} else {
		return "#f"
	}
}

// IsBool implements Term.
func (Bool) IsBool() bool {
	return true
}

// IOsNumber implements Term.
func (Bool) IsNumber() bool {
	return false
}

// IsPair implements Term.
func (Bool) IsPair() bool {
	return false
}

// IsString implements Term.
func (Bool) IsString() bool {
	return false
}

// IsSymbol implements Term.
func (Bool) IsSymbol() bool {
	return false
}
