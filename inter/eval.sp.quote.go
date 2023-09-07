package inter

func init() {
	Register("quote",
		false, // no args evaluation !
		func(_ *Inter, args Term) Term {
			return args // return as is
		})
}
