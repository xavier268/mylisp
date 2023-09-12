package inter

import (
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
