package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xavier268/mylisp/mytest"
)

func TestNumberScan(t *testing.T) {
	var tests = []string{
		"1",
		"-0",
		"-50",
		"010",
		"-5/6",
		"14/7",
		"14/-7",
		"12.",
		"12.3",
	}

	sb := new(strings.Builder)
	for i, n := range tests {
		N, err := NumberFromString(n)
		fmt.Fprintf(sb, "%2d:\t%s\t --> %s\n", i, n, N)
		if err != nil {
			fmt.Fprintf(sb, "\t*******  %v\n", err)
		}
	}

	mytest.Verify(t, sb.String(), "number.scan")

}
