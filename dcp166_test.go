package daily_coding_practice

import "testing"

type testCase166 struct {
	Name   string
	Input  [][]int
	Output []int
}

var testCases166 = []testCase166{
	{
		Name:   "empty array",
		Input:  nil,
		Output: nil,
	},
}

func Test166(t *testing.T) {
	for _, val := range testCases166 {
		t.Run(val.Name, func(t *testing.T) {
			runTest166(val, t)
		})
	}
}

func runTest166(testCase testCase166, t *testing.T) {
	//
}
