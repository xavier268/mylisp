package inter

import "fmt"

var ErrNaN = fmt.Errorf("term is not a number")

func init() {
	Register("+", true, spPlus, []string{"addition for numbers"})
	Register("*", true, spTimes, []string{"multiplication for numbers"})
	Register("/", true, spDiv, []string{"division for numbers"})
	Register("-", true, spMinus, []string{"subtraction for numbers"})
}

// ( + ) - 0
// ( + n ) - n
// ( + n m ) - n + m
// ( + n m p ) - n + m + p
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

// ( - ) -> 0
// ( - n ) -> -n
// ( - n m ) -> n-m
// ( - n m p q ) ->
func spMinus(it *Inter, args Term) Term {

	n := car(args)

	if n == nil || (n == Pair{}) { // ( + ) evaluates to 0
		return NumberZero
	}

	if !n.IsNumber() {
		return Error{ErrNaN, args}
	}

	if cdr(args) == nil {
		return NumberZero.Minus(n.(Number)) // -n
	}

	m := spPlus(it, cdr(args)) // not a mistake, because - is associative left to right
	if m.IsNumber() {
		return n.(Number).Minus(m.(Number))
	} else {
		return Error{ErrNaN, args}
	}
}

// ( * ) - 1
// ( * n ) - n
// ( * n m ) - nm
func spTimes(it *Inter, args Term) Term {

	n := car(args)
	if n == nil || (n == Pair{}) {
		return NumberOne
	}

	if !n.IsNumber() {
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

// ( / n ) - 1/n
// ( / n m ) - n/m
// ( / n  m p q ) - n/mpq
// ( / ) - 1
func spDiv(it *Inter, args Term) Term {

	n := car(args)
	if n == nil || (n == Pair{}) {
		return NumberOne
	}
	if !n.IsNumber() {
		return Error{ErrNaN, args}
	}

	if cdr(args) == nil {
		return NumberOne.Div(n.(Number))
	}

	m := spTimes(it, cdr(args)) // Not an error, because / associates from LEFT to RIGHT
	if m.IsNumber() {
		return n.(Number).Div(m.(Number))
	} else {
		return Error{ErrNaN, args}
	}
}
