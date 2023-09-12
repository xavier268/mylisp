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
}

func spHelp(_ *Inter, t Term) Term { // ( help) for general help or ( help keyword ) for help on keyword - keyword is a symbol

	// t contains ( kw1 kw2 ...)
	sb := new(strings.Builder)

	fmt.Fprintf(sb, "; available keywords : %s", spKeywords(nil, nil))

	for rest := t; rest != nil; rest = cdr(rest) {

		kw := car(rest)
		if kw != nil {
			if sym, ok := kw.(Symbol); ok {
				if kwdata, ok := KEYWORDS[sym.Value]; ok {
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
