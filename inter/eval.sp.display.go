package inter

import "fmt"

func init() {
	Register("display", true, spDisplay, []string{
		"(display x )",
		"Evaluates a single Term x and displays it on the terminal.",
		"x can be any Term.",
		"There will be no space before nor after.",
		"If x is a String, it will be displayed unquoted. That will not apply for Strings in other elements, like Pairs.",
	})
	Register("newline", false, spNewline, []string{
		"returns a newline char as a string. Takes no arguments.",
		"If arguments are provided, they are not evaluated and silently ignored.",
	})
	Register("debug", true, spDebug, []string{
		"( debug value ) or ( debug )",
		"set the debug flag to an integer value, and return its new value.",
		" if no argument is provided, just return the current value",
	})
}

var ErrDisplayTooManyArguments = fmt.Errorf("display: too many arguments")

// ( display x ) will display a SINGLE Term, x.
// x can be any Term.
// There will be no space before nor after.
// If x is a Sring, it will be displayed unquoted. That will not apply for Strings in other elements, like Pairs.
// spDisplay is called with (x).
func spDisplay(_ *Inter, t Term) Term {
	if car(t).IsString() { // special de-quoting of strings
		fmt.Print(car(t).(String).Value)
	} else {
		fmt.Printf("%s", ToString(car(t)))
	}
	if cdr(t) != nil {
		return Error{ErrDisplayTooManyArguments, t}
	}
	return nil
}

// ( newline ) // with ZERO arguments, prints a new line.
func spNewline(_ *Inter, t Term) Term {
	fmt.Println()
	return nil
}

func spDebug(_ *Inter, t Term) Term { // t contains ( value ) or ()
	if t == nil || (t == Pair{}) {
		return Number{DEBUG, 1}
	}

	if car(t) != nil {
		if n, ok := car(t).(Number); ok && n.Den == 1 {
			DEBUG = n.Num
			return n
		}
	}
	var ErrNotAnInteger = fmt.Errorf("not an integer")
	return Error{ErrNotAnInteger, t}
}
