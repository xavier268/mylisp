package inter

import (
	"cmp"
	"slices"
)

func init() {
	Register("keywords", false, spKeywords, []string{"displays a list of primitive keywords"})
}

// ( keywords ... ) will return the current list of built-in keywords currently implemented.
// Arguments will simply be ignored and never evaluated.
// Keywords are sorted by alphabetic order.
func spKeywords(*Inter, Term) Term {

	var kw Term
	kws := make([]string, 0, len(keywordsDefinitions))

	for s := range keywordsDefinitions {
		kws = append(kws, s)
	}

	// internally, sort in reverse order
	slices.SortFunc(kws, func(a, b string) int { return cmp.Compare(b, a) })

	for _, s := range kws {
		kw = Pair{
			Car: Symbol{s},
			Cdr: kw,
		}
	}

	return kw
}
