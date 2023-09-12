package inter

func init() {
	Register("quote",
		false, // no args evaluation !
		func(_ *Inter, args Term) Term {
			return args // return as is
		}, []string{
			" ' a , or ( quote . a )",
			" ' ( a b ...) or ( quote a b ... ) ",
			" returns the symbol or quoted list without evaluating it.",
		})
}
