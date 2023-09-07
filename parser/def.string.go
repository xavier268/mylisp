package parser

import "fmt"

var _ Term = String{}

type String struct {
	Value string
}

// IsBool implements Term.
func (String) IsBool() bool {
	return false
}

// IsNumber implements Term.
func (String) IsNumber() bool {
	return false
}

// IsPair implements Term.
func (String) IsPair() bool {
	return false
}

// IsString implements Term.
func (String) IsString() bool {
	return true
}

// IsSymbol implements Term.
func (String) IsSymbol() bool {
	return false
}

func (t String) String() string {
	return fmt.Sprintf("%q", t.Value)
}
