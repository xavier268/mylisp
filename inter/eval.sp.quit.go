package inter

import (
	"fmt"
	"os"
)

// built in form that quit immediately.
// Syntax : ( quit ... )
func spQuit(it *Inter, args Term) Term {

	if it.OnQuit != nil {
		it.OnQuit()
	}
	fmt.Println("Bye.")

	os.Exit(0) // actual quit
	return nil // keep the compiler happy ...
}

// register quit function.
func init() {
	Register("quit", false, spQuit, []string{"exit the interpreter"})
}
