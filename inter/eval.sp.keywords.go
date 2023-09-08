package inter

func init() {
	Register("keywords", false, spKeywords)
}

// ( keywords ... ) will return the current list of built-in keywords currently implemented.
// Arguments will simplu be ignored and never evaluated.
func spKeywords(*Inter, Term) Term {

	var kw Term
	for s := range KEYWORDS {
		kw = Pair{
			Car: Symbol{s},
			Cdr: kw,
		}
	}

	return kw
}
