// package inter constains lexer, parser and interpreter.
package inter

//go:generate go get golang.org/x/tools/cmd/goyacc

//go:generate go install golang.org/x/tools/cmd/goyacc

//go:generate goyacc -o parser.go -p "my" grammar.y

//go:generate go mod tidy

//go:generate go fmt ./...

const (
	VERSION = "0.1.0"
)
