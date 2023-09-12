package examples

import (
	"testing"

	"github.com/xavier268/mylisp/inter"
)

func TestExamples(t *testing.T) {

	it := inter.NewInter()
	t1, err := inter.Parse1String(`( load "demo1.scm")`, "demo1.scm")
	if err != nil {
		t.Fatalf("unexpected parse error : %v", err)
	}
	t2 := it.Eval(t1)
	if t2 != nil && t2.IsError() {
		t.Fatalf("error on loading : %v", inter.ToString(t2))
	}
}
