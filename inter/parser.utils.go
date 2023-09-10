package inter

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Return the LAST parsed Term.
func Parse1(rdr io.Reader, sourcename string) (Term, error) {
	r, err := ParseN(rdr, sourcename)
	if len(r) == 0 {
		return nil, err
	} else {
		return r[len(r)-1], err
	}
}

// Return a (golang) list of parsed Terms.
func ParseN(rdr io.Reader, sourcename string) ([]Term, error) {
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

func Parse1File(filename string) (Term, error) {
	var err error
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Parse1(f, filename)
}

func ParseNFile(filename string) ([]Term, error) {
	var err error
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseN(f, filename)
}

func Parse1String(input string, sourcename string) (Term, error) {
	return Parse1(strings.NewReader(input), sourcename)
}

func ParseNString(input string, sourcename string) ([]Term, error) {
	return ParseN(strings.NewReader(input), sourcename)
}

func MustParse1String(input string, sourcename string) Term {
	r, err := Parse1String(input, sourcename)
	if err != nil {
		panic(err)
	}
	return r
}

func MustParseNString(input string, sourcename string) []Term {
	r, err := ParseNString(input, sourcename)
	if err != nil {
		panic(err)
	}
	return r
}
