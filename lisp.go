package main

import (
	"fmt"

	"github.com/xavier268/mylisp/inter"
)

func main() {

	inter.DEBUG = 0

	fmt.Print(inter.Welcome())

	it := inter.NewInter()

	it.Repl()

	fmt.Println("\nBye!")

}
