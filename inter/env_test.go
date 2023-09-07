package inter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xavier268/mylisp/mytest"
	"github.com/xavier268/mylisp/parser"
)

func TestBind(t *testing.T) {
	tests := [][]string{
		{"(X Y)", "(a b c)"},
		{"(X Y)", "(a b)"},
		{"(X Y)", "(a )"},
		{"(X Y)", "()"},
		{"(X Y)", "( . ())"},
		{"(X Y)", ""},

		{" ( A . B)", "(a (b c) d)"},
		{" ( A . B)", "(a (b ))"},
		{" ( A . B)", "(a b)"},
		{" ( A . B)", "(a )"},
		{" ( A . B)", "( )"},
		{" ( A . B)", ""},

		{"S", "(a b)"},
		{"S", "(a )"},
		{"S", "()"},
		{"S", "( . ())"},
		{"S", ""},

		{"(X ( Y Z ) U)", "(a (b c) d)"},
		{"(X ( Y 3 ) U)", "(a (b c) d)"},

		{"(X ( Y Z ) U)", "(a (b) d)"},
		{"(X ( Y Z ) U)", "(a b d)"},
	}

	sb := new(strings.Builder)

	for i, s := range tests {

		fmt.Fprintln(sb)
		it := NewInter()

		s0, err := parser.ParseString(s[0], t.Name())

		if err != nil {
			fmt.Fprintln(sb, "*******", err, s[0])
			continue
		}
		s1, err := parser.ParseString(s[1], t.Name())
		if err != nil {
			fmt.Fprintln(sb, "*******", err, s[1])
			continue
		}
		fmt.Fprintf(sb, "%d: formal args %s\t target %s\n", i, s0, s1)

		it.Bind(s0, s1)
		fmt.Fprintln(sb, it)
	}

	mytest.Verify(t, sb.String(), "env.bind")

}
