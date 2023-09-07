package parser

import "fmt"

var ErrNotAList = fmt.Errorf("not a list")

var _ Term = Pair{}

type Pair struct {
	Car, Cdr Term
}

// IsProcedure implements Term.
func (Pair) IsProcedure() bool {
	return false
}

// IsBool implements Term.
func (Pair) IsBool() bool {
	return false
}

// IsNumber implements Term.
func (Pair) IsNumber() bool {
	return false
}

// IsPair implements Term.
func (Pair) IsPair() bool {
	return true
}

// IsString implements Term.
func (Pair) IsString() bool {
	return false
}

// IsSymbol implements Term.
func (Pair) IsSymbol() bool {
	return false
}

func (t Pair) String() string {

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

// if c is a list, returns the string of its inside, without parenthesis.
// if not, return error.
func (c *Pair) isList() (s string, err error) {
	if c == nil {
		return "", ErrNotAList
	}
	if (*c == Pair{}) {
		return "", ErrNotAList
	}
	if c.Cdr == nil {
		return c.Car.String(), nil
	}
	if cc, ok := c.Cdr.(Pair); ok {
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
