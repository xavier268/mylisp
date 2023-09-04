package parser

import (
	"cmp"
	"fmt"
	"strconv"
)

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

// a convenience string function, thaht handles nil Term in the correct way.
func ToString(t Term) string {
	if t == nil {
		return "<nil>"
	}
	return t.String()
}

func (t Cell) String() string {

	if t.Car == nil { //car == nil
		if t.Cdr == nil {
			return "( )"
		} else {
			return "( . " + t.Cdr.String() + " )"
		}
	} else { //car != nil
		ss, err := t.isList()
		if err == nil {
			return "( " + ss + " )"
		} else {
			return "( " + t.Car.String() + " . " + t.Cdr.String() + " )"
		}

	}

}

var ErrNotAList = fmt.Errorf("not a list")

// if c is a list, returns the string of its inside, without parenthesis.
// if not, retun error.
func (c *Cell) isList() (s string, err error) {
	if c == nil {
		return "", ErrNotAList
	}
	if (*c == Cell{}) {
		return "", nil
	}
	if c.Cdr == nil {
		return c.Car.String(), nil
	}
	if cc, ok := c.Cdr.(Cell); ok {
		// looks like a list
		scc, err := cc.isList()
		if err != nil {
			return "", err
		} else {
			return c.Car.String() + " " + scc, nil
		}
	} else {
		return "", ErrNotAList
	}
}

// numbers are rational numbers
// Num & Den are always prime among them,
// Num is always > 0, except to represent NaN with 1/0.
type Number struct {
	Num int
	Den int
}

var NumberZero = Number{0, 1}
var NumberOne = Number{1, 1}
var NumberNaN = Number{1, 0}

func (n Number) String() string {
	if n.Den == 0 {
		return "NaN"
	}
	if n.Den == 1 {
		return fmt.Sprintf("%d", n.Num)
	}
	return fmt.Sprintf("%d/%d", n.Num, n.Den)
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

func NumberFromString(s string) (n Number, err error) {

	if pp := PAT_NUMBER.FindStringIndex(s); pp == nil || pp[0] != 0 || pp[1] != len(s) {
		return NumberZero, fmt.Errorf("invalid number format : <%s>", s)
	}
	m := PAT_NUMBER.FindStringSubmatch(s)
	if len(m) != 5 {
		return NumberZero, fmt.Errorf("invalid number format : <%s>", s)
	}

	n.Num, err = strconv.Atoi(m[2])
	n.Den = 1
	if err != nil {
		return NumberZero, fmt.Errorf("invalid number format : <%s> because %v", s, err)
	}
	if len(m[1]) > 0 { // minus sign
		n.Num = -n.Num
	}
	if len(m[4]) > 0 { // denominator
		n.Den, err = strconv.Atoi(m[4])
		if err != nil {
			return NumberZero, fmt.Errorf("invalid number format : <%s> because %v", s, err)
		}
	}
	return n.Normalize(), nil
}

func (n Number) Plus(m Number) Number {
	return Number{
		Num: n.Num*m.Den + m.Num*n.Den,
		Den: n.Den * m.Den,
	}.Normalize()
}
func (n Number) Minus(m Number) Number {
	return Number{
		Num: n.Num*m.Den - m.Num*n.Den,
		Den: n.Den * m.Den,
	}.Normalize()
}

func (n Number) Mul(m Number) Number {
	return Number{
		Num: n.Num * m.Num,
		Den: n.Den * m.Den,
	}.Normalize()
}

func (n Number) Div(m Number) Number {
	return Number{
		Num: n.Num * m.Den,
		Den: n.Den * m.Num,
	}.Normalize()
}

// Compare n and m. -1 if n<m, 0 if n==m, 1 if n>m.
func Compare(n, m Number) int {
	return cmp.Compare(n.Num*m.Den, m.Num*n.Den)
}

// ===================================================================

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
