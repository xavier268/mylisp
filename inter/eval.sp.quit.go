package inter

// built in form that quit.
// Syntax : ( quit ... )
func spQuit(it *Inter, args Term) Term {
	it.quit = true
	return nil
}

// register quit function.
func init() {
	Register("quit", false, spQuit)
}
