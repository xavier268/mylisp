package inter

// Error is a type to indicate an error has occured.
type Error struct {
	Err     error
	Context Term
}

// IsError implements Term.
func (Error) IsError() bool {
	return true
}

var _ Term = Error{}

// IsBool implements Term.
func (Error) IsBool() bool {
	return false
}

// IsNumber implements Term.
func (Error) IsNumber() bool {
	return false
}

// IsPair implements Term.
func (Error) IsPair() bool {
	return false
}

// IsProcedure implements Term.
func (Error) IsProcedure() bool {
	return false
}

// IsString implements Term.
func (Error) IsString() bool {
	return false
}

// IsSymbol implements Term.
func (Error) IsSymbol() bool {
	return false
}

func (e Error) String() string {
	return e.Err.Error() + " in " + e.Context.String()
}
