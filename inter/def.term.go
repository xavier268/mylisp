package inter

import (
	"fmt"
	"strings"
)

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
func ToString(t any) string {
	if t == nil {
		return "<nil>"
	}
	switch tt := t.(type) {
	case Term:
		return tt.String()
	case []Term:
		res := []string{}
		for _, rt := range tt {
			res = append(res, rt.String())
		}
		return strings.Join(res, "\n")
	default:
		return fmt.Sprintf("%v", t)

	}
}
