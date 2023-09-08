package inter

import "fmt"

func init() {
	Register("display", true, spDisplay)
	Register("newline", false, spNewline)
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
	if t != nil || (t != Pair{}) {
		return Error{ErrDisplayTooManyArguments, t}
	}
	return nil
}
