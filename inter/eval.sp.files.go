package inter

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

func init() {
	Register("join-path", true, spJoinPath, []string{
		"( join-path p1 p2 p3 ...)",
		"evaluate the arguments, and join them into a path, using the file separator for the system.",
	})
	Register("file-sep", false, spFileSep, []string{
		"returns the system file separator as a string",
	})
	Register("os-name", false, spOSName, []string{
		"returns the system os name as a string, eg windows-amd64, from GOOS and GOARCH",
	})
	Register("string-append", true, spStringAppend, []string{
		"evaluate the arguments, and join them into a single string.",
	})
	Register("load", true, spLoad, []string{
		"When you use load, it reads and evaluates the code in the specified file,",
		"just as if you had typed the code directly into the REPL or program.",
		"This means that any definitions, including variable bindings and function definitions, ",
		"become available in your current Scheme environment.",
		"load is typically used for loading complete Scheme programs or libraries. ",
		"It's useful when you want to use functions or variables defined in another file as part of your program. ",
		"For example, if you have a utility library that you want to reuse across multiple projects, you can load it using load.",
		"load can create or modify variables and procedures in the current environment, so it can affect the global scope of your program. ",
		"This can lead to potential naming conflicts if you're not careful.",
		"load will only return the last evaluated value.",
	})
}

func spJoinPath(it *Inter, t Term) Term { // t contains the path components ( p1 p2 p3 ... )

	if t == nil || t == (Pair{}) {
		return String{}
	}

	res := make([]string, 0, 6)

	for rest := t; rest != nil && (rest != Pair{}); rest = cdr(rest) {
		if car(t) != nil {
			switch tt := car(rest).(type) {
			case String:
				res = append(res, tt.Value) // do not quote the strings !
			default:
				res = append(res, car(rest).String())
			}
		}
	}

	return String{filepath.Join(res...)}
}

func spStringAppend(it *Inter, t Term) Term { // t contains the substrings components ( s1 s2 s3 ... ) that will be concatenated.
	if t == nil || t == (Pair{}) {
		return String{}
	}

	res := make([]string, 0, 6)

	for rest := t; rest != nil && (rest != Pair{}); rest = cdr(rest) {
		if car(t) != nil {
			switch tt := car(rest).(type) {
			case String:
				res = append(res, tt.Value) // do not quote the strings !
			default:
				res = append(res, car(rest).String())
			}
		}
	}
	return String{strings.Join(res, "")}
}

func spFileSep(_ *Inter, _ Term) Term { // returns the system file separator as a string
	return String{string(filepath.Separator)}
}

func spOSName(_ *Inter, _ Term) Term { // returns the system os name as a string, eg windows-amd64, from GOOS and GOARCH
	return String{runtime.GOOS + "-" + runtime.GOARCH}
}

var ErrLoadArgument = fmt.Errorf("load expects a single, String, argument")

func spLoad(it *Inter, t Term) Term {

	if t == nil || car(t) == nil || !t.IsPair() || !car(t).IsString() || cdr(t) != nil {
		return Error{ErrLoadArgument, t}
	}

	filename := car(t).(String).Value
	terms, err := ParseNFile(filename)
	if err != nil {
		return Error{err, t}
	}

	var res Term // only store and return the LAST value.
	for _, term := range terms {
		res = it.Eval(term)
	}

	return res
}
