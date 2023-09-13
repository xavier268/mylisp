package inter

import "fmt"

func init() {
	Register("display", true, spDisplay, []string{
		"(display x1 x2 ... )",
		"Evaluates Terms xi and displays them on the terminal.",
		"xi can be any Term.",
		"There will be no space before nor after each xi.",
		"If x is a String, it will be displayed unquoted. That will not apply for Strings in other elements, like Pairs.",
		"always return nil",
	})
	Register("newline", false, spNewline, []string{
		"returns a newline char as a string. Takes no arguments.",
		"If arguments are provided, they are not evaluated and silently ignored.",
	})

}

var ErrDisplayTooManyArguments = fmt.Errorf("display: too many arguments")

// x can be any Term.
// There will be no space before nor after.
// If x is a String, it will be displayed unquoted. That will not apply for Strings within other elements, like Pairs.
// spDisplay is called with (x).
// always return nil
func spDisplay(_ *Inter, t Term) Term { // t contains ( x1 x2 ...)

	for rest := t; rest != nil && (rest != Pair{}); rest = cdr(rest) {
		if car(rest) != nil { // just ignore and skip nil values
			if car(rest).IsString() { // special de-quoting of strings
				fmt.Print(car(rest).(String).Value)
			} else {
				fmt.Printf("%s", ToString(car(rest)))
			}
		}
	}
	return nil
}

// ( newline ) // with ZERO arguments, prints a new line.
func spNewline(_ *Inter, t Term) Term {
	fmt.Println()
	return nil
}
