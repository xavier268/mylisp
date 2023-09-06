package inter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xavier268/mylisp/mytest"
)

func TestMakeList(t *testing.T) {

	X, Y, Z := Symbol{Value: "x"}, Symbol{Value: "y"}, Symbol{Value: "z"}

	tests := [][]Term{

		{X},
		{X, makeList(Y, Z), X},
		{},
	}
	sb := new(strings.Builder)
	for i, l := range tests {
		fmt.Fprintln(sb)
		if len(l) == 0 {
			fmt.Fprintf(sb, "%d:\t%v ---> %v\n", i, l, makeList())
			fmt.Fprintf(sb, "%d:\t%v ---> %#v\n", i, l, makeList())
		} else {
			fmt.Fprintf(sb, "%d:\t%v ---> %v\n", i, l, makeList(l...))
			fmt.Fprintf(sb, "%d:\t%v ---> %#v\n", i, l, makeList(l...))
		}
	}

	mytest.Verify(t, sb.String(), "makelist")
}
