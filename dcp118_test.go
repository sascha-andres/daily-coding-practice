package daily_coding_practice

import (
	"github.com/sascha-andres/dcp114/utl"
	"testing"
)

type testCase118 struct {
	input  []int
	output []int
}

var testCases118 = []testCase118{
	{input: []int{-9, -2, 0, 2, 3}, output: []int{0, 4, 4, 9, 81}},
	{input: []int{}, output: []int{}},
	{input: []int{0}, output: []int{0}},
}

func Test118(t *testing.T) {
	for _, value := range testCases118 {
		runTest118(value.input, value.output, t)
	}
}

/////////////////////////////////////////////////////////////////

func runTest118(input, output []int, t *testing.T) {
	result := Run118(input)
	if len(output) != len(result) || !utl.ArrayMatchInt(result, output) {
		t.Logf("for %v got %v expected %v", input, result, output)
		t.Fail()
	}
}
