package main

import (
	"flag"
	"fmt"

	"github.com/xavier268/mylisp/inter"
)

var flag_version, flag_help bool
var loadFile string

func init() {
	flag.IntVar(&inter.DEBUG, "d", 0, "debug level 0: none, 1 : light, 2: trace")
	flag.BoolVar(&flag_version, "v", false, "version")
	flag.BoolVar(&flag_help, "h", false, "help")
	flag.StringVar(&loadFile, "l", "", "load file")

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

	// create interpreter
	it := inter.NewInter()

	// preload file if any specified
	if loadFile != "" {
		fmt.Println("loading file : ", loadFile)
		terms, err := inter.ParseNFile(loadFile)
		if err != nil {
			fmt.Println(err)
		} else {
			for _, term := range terms {
				t := it.Eval(term)
				if t != nil && t.IsError() {
					fmt.Println(t)
				}
			}
		}
	}

	//  repl loop
	it.Repl()

	fmt.Println("\nBye!")
}
