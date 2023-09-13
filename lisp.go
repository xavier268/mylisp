package main

import (
	"flag"
	"fmt"

	"github.com/xavier268/mylisp/inter"
)

var flag_version, flag_help bool

func init() {
	flag.IntVar(&inter.DEBUG, "d", 0, "debug level 0: none, 1 : light, 2: trace")
	flag.BoolVar(&flag_version, "v", false, "version")
	flag.BoolVar(&flag_help, "h", false, "help")
}

func main() {

	// Display version and welcome message
	fmt.Println(inter.Welcome())

	// handle flags()
	flag.Parse()
	if flag_version {
		return
	}

	if flag_help {
		flag.PrintDefaults()
		return
	}

	// actual repl

	fmt.Print(inter.Welcome())

	it := inter.NewInter()

	it.Repl()

	fmt.Println("\nBye!")

}
