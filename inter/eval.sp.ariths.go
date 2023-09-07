package inter

import "fmt"

var ErrNaN = fmt.Errorf("term is not a number")

func init() {
	Register("+", true, spPlus)
	Register("*", true, spTimes)
}

func spPlus(it *Inter, args Term) Term {

	n := car(args)

	if n == nil || (n == Pair{}) { // ( + ) evaluates to 0
		return NumberZero
	}

	if !n.IsNumber() {
		return Error{ErrNaN, args}
	}

	if cdr(args) == nil {
		return n
	}

	m := spPlus(it, cdr(args))
	if m.IsNumber() {
		return n.(Number).Plus(m.(Number))
	} else {
		return Error{ErrNaN, args}
	}
}

func spTimes(it *Inter, args Term) Term {

	n := car(args)
	if n == nil || !n.IsNumber() {
		return Error{ErrNaN, args}
	}

	if cdr(args) == nil {
		return n
	}

	m := spTimes(it, cdr(args))
	if m.IsNumber() {
		return n.(Number).Mul(m.(Number))
	} else {
		return Error{ErrNaN, args}
	}
}
