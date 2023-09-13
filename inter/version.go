// package inter constains lexer, parser and interpreter.
package inter

import "fmt"

//go:generate go get golang.org/x/tools/cmd/goyacc

//go:generate go install golang.org/x/tools/cmd/goyacc

//go:generate goyacc -o parser.go -p "my" grammar.y

//go:generate go mod tidy

//go:generate go fmt ./...

const (
	VERSION   = "0.2.3"
	COPYRIGHT = "(c) Xavier Gandillot (aka xavier268) 2023"
)

// Prinţ welcome message.
func Welcome() string {
	return fmt.Sprintf("Welcome to mylisp, a minimal lisp/scheme interpreter\nVersion %s - %s\n", VERSION, COPYRIGHT)
}

// 0 - off, 1 - on, 2 - on with trace
var DEBUG int = 0
