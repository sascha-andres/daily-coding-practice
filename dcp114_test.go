package daily_coding_practice

import (
	"testing"
)

var testCases114 = [][]string{
	{"hello", "hello", ":/"},
	{"hello/world:here", "here/world:hello", ":/"},
	{"hello/world:here/", "here/world:hello/", ":/"},
	{"hello//world:here", "here/world/:hello", ":/"},
	{"//", "//", ":/"},
	{"", "", ":/"},
}

func Test114(t *testing.T) {
	for _, val := range testCases114 {
		runTest114(val[0], val[2], val[1], t)
	}
}

/////////////////////////////////////////////////////////////////

func runTest114(input string, separators string, output string, t *testing.T) {
	result := Run(input, separators)
	if result != output {
		t.Logf("got '%s', expected '%s' for '%s' with separators '%s'", result, output, input, separators)
		t.Fail()
	}
}
