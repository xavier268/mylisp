package main

import (
	"fmt"

	"github.com/xavier268/mylisp/inter"
)

func main() {

	inter.DEBUG = 0

	inter.Welcome()

	it := inter.NewInter()

	it.Repl()

	fmt.Println("\nBye!")

}
