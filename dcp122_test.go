package daily_coding_practice

import "testing"

type testCase122 struct {
	Matrix   [][]int
	Expected int
}

var testCases122 = []testCase122{
	{[][]int{
		{0, 3, 1, 1},
		{2, 0, 0, 4},
		{1, 5, 3, 1},
	}, 12},
	{[][]int{
		{1, 3, 1, 1},
		{2, 0, 0, 4},
		{1, 5, 3, 1},
	}, 13},
	{nil,
		0},
	{[][]int{nil},
		0}, {[][]int{
		{0, 3, 1, 1},
		{2, 0, 0, 4, 2},
		{1, 5, 3, 1},
	}, 0},
}

func Test122(t *testing.T) {
	for _, val := range testCases122 {
		runTest122(val.Matrix, t, val.Expected)
	}
}

func runTest122(testCase [][]int, t *testing.T, expected int) {
	result := Run122(testCase)
	if result != expected {
		t.Logf("expected %d got %d", expected, result)
		t.Fail()
	}
}
