package parser

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Parse(rdr io.Reader, sourcename string) (Term, error) {
	var err error
	lx := NewLexer(rdr, sourcename)
	p := myNewParser()
	e := p.Parse(lx)
	if e != 0 {
		err = fmt.Errorf("error : %v", lx.LastErr)
	}
	r := lx.LastResult
	return r, err
}

func ParseFile(filename string) (Term, error) {
	var err error
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Parse(f, filename)
}

func ParseString(input string, sourcename string) (Term, error) {
	return Parse(strings.NewReader(input), sourcename)
}

func MustParseString(input string, sourcename string) Term {
	r, err := ParseString(input, sourcename)
	if err != nil {
		panic(err)
	}
	return r
}
