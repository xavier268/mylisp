package inter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xavier268/mylisp/mytest"
)

func TestMakeList(t *testing.T) {

	X, Y, Z := Atom{Value: "x"}, Atom{Value: "y"}, Atom{Value: "z"}

	tests := [][]Term{

		{X},
		{X, MakeList(Y, Z), X},
		{},
	}
	sb := new(strings.Builder)
	for i, l := range tests {
		fmt.Fprintln(sb)
		if len(l) == 0 {
			fmt.Fprintf(sb, "%d:\t%v ---> %v\n", i, l, MakeList())
			fmt.Fprintf(sb, "%d:\t%v ---> %#v\n", i, l, MakeList())
		} else {
			fmt.Fprintf(sb, "%d:\t%v ---> %v\n", i, l, MakeList(l...))
			fmt.Fprintf(sb, "%d:\t%v ---> %#v\n", i, l, MakeList(l...))
		}
	}

	mytest.Verify(t, sb.String(), "makelist")
}
