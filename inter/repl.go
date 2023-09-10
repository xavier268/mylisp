package inter

import (
	"bufio"
	"fmt"
	"os"
)

func (it *Inter) Repl() {

	for {
		lines := readLines()

		tl, err := ParseNString(lines, "terminal")
		if err != nil {
			fmt.Println("Error :", err)
			continue
		}
		for _, t := range tl {
			tt := it.Eval(t)
			fmt.Printf("%s\n", ToString(tt))
		}
	}
}

// basic interactive entry.
// reads until parenthesis count matches.
func readLines() (lines string) {

	scanner := bufio.NewScanner(os.Stdin)
	lines = ""

	for {
		fmt.Printf("> ")
		scanner.Scan()
		lines = lines + "\n" + scanner.Text()

		done, err := finished(lines)
		if err != nil {
			fmt.Println("Error :", err)
			lines = ""
			continue
		}
		if done {
			return lines
		}
	}
}

var ErrParenthesisUnbalanced = fmt.Errorf("too many closing parenthesis")

// if form looks correct, return true,nil.
// returns error if form can no longer be correct.
// return false, nil to indicate continue entry.
func finished(lines string) (bool, error) {

	left, right := 0, 0
	comment := false
	quote := false
	for _, s := range lines {
		switch s {
		case '(':
			if !quote && !comment {
				left += 1
			}
		case ')':
			if !quote && !comment {
				right += 1
			}
		case '\n':
			comment = false
		case '"':
			if !comment {
				quote = !quote
			}
		case ';':
			comment = true
		}
		// realtime checks
		if right > left {
			return false, ErrParenthesisUnbalanced
		}
	}
	// end of string checks
	if quote {
		return false, nil
	}
	if left == right {
		return true, nil
	}
	return false, nil
}
