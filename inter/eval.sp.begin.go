package inter

func init() {
	Register("begin", true, spBegin, []string{
		"(begin expr1 expr2 expr3 ...)",
		" begin will ensure the expressions are sequentially evaluated,",
		"in the current environnement, returning the value of the last one.",
	})
}

func spBegin(it *Inter, t Term) Term { // called with ( expr1 expr2 expr3 ...)

	var res Term
	for rest := t; rest != nil && (rest != Pair{}); rest = cdr(rest) {
		res = car(rest)
	}
	return res
}
