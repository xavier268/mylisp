package inter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xavier268/mylisp/mytest"
)

func TestMakeList(t *testing.T) {

	X, Y, Z := Atom{"x"}, Atom{"y"}, Atom{"z"}

	tests := [][]Term{
		[]Term{},
		{X},
		{X, MakeList(Y, Z), X},
	}
	sb := new(strings.Builder)
	for i, l := range tests {
		fmt.Fprintln(sb)
		fmt.Fprintln(sb, i, l, "---->", MakeList(l...))
	}

	mytest.Verify(t, sb.String(), "makelist")
}
