package inter

import (
	"fmt"
	"strings"
)

func init() {
	// ( help) or ( help keyword1 keyword2 ...)
	// keywords must be symbols, and are never evaluated.
	Register("help", false, spHelp, []string{
		"( help symbol1 symbol2 ... ) or ( help )",
		"displays help for the provided symbols, if available.",
		"returns nil",
	})

	Register("version", false, func(_ *Inter, _ Term) Term { return String{VERSION} }, []string{
		"( version )",
		"returns the version of the interpreter as a String",
		"arguments, if any, are not evaluated and ignored",
	})

	Register("debug", true, spDebug, []string{
		"( debug value ) or ( debug )",
		"set the debug flag to an integer value, and return its new value.",
		" if no argument is provided, just return the current value",
	})
}

func spHelp(_ *Inter, t Term) Term { // ( help) for general help or ( help keyword ) for help on keyword - keyword is a symbol

	// t contains ( kw1 kw2 ...)
	sb := new(strings.Builder)

	fmt.Fprintf(sb, "; available keywords : %s", spKeywords(nil, nil))

	for rest := t; rest != nil; rest = cdr(rest) {

		kw := car(rest)
		if kw != nil {
			if sym, ok := kw.(Symbol); ok {
				if kwdata, ok := keywordsDefinitions[sym.Value]; ok {
					fmt.Fprintf(sb, "\n;\n; help for %s\n", sym.Value)
					for _, h := range kwdata.help {
						fmt.Fprintf(sb, "; %s\n", h)
					}
				}
			}
		}
	}

	fmt.Println(sb.String())
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
