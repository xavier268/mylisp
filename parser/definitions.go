package parser

import "fmt"

type Term interface {
	String() string
}

var _ Term = Cell{}
var _ Term = Atom{}
var _ Term = String{}
var _ Term = Number{}

type Cell struct {
	Car, Cdr Term
}

func (t Cell) String() string {

	panic("not implemented")
}

// numbers are rational numbers
// Num & Den are always prime among them,
// Num is always > 0, except to represent NaN with 1/0.
type Number struct {
	Num int
	Den int
}

func (t Number) String() string {
	return fmt.Sprintf("%d/%d", t.Num, t.Den)
}

// Normalize the internal representation of a number.
// 0/x is normalized as 0/1,
// NaN is normalized 1/0
// Den is always > 0
func (n Number) Normalize() Number {

	if n.Num == 0 {
		return Number{
			Num: 0,
			Den: 1,
		}
	}

	if n.Den == 0 {
		return Number{
			Num: 1,
			Den: 0,
		}
	}

	p := Gcd(n.Num, n.Den)
	if n.Den > 0 {
		return Number{
			Num: n.Num / p,
			Den: n.Den / p,
		}
	}
	if n.Den < 0 {
		return Number{
			Num: -n.Num / p,
			Den: -n.Den / p,
		}
	}
	panic("code should be unreacheable")
}

type String struct {
	Value string
}

func (t String) String() string {
	return fmt.Sprintf("%q", t.Value)
}

type Atom struct {
	Value string
}

func (t Atom) String() string {
	return t.Value
}
