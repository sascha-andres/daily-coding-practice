package daily_coding_practice

import "testing"

type testCase165 struct {
	Input    []int
	Output   []int
	HasError bool
}

var testCases165 = []testCase165{
	{
		Input:    nil,
		Output:   nil,
		HasError: true,
	},
	{
		Input:    []int{3, 4, 9, 6, 1},
		Output:   []int{1, 1, 2, 1, 0},
		HasError: false,
	},
}

func Test165(t *testing.T) {
	for _, val := range testCases165 {
		runTest165(val, t)
	}
}

func runTest165(test testCase165, t *testing.T) {
	actual, err := Run165(test.Input)
	if test.HasError && err == nil {
		t.Logf("expected error, received non")
		t.Fail()
		return
	}
	if !test.HasError && err != nil {
		t.Logf("got error: %v", err)
		t.Fail()
		return
	}
	if len(actual) != len(test.Output) {
		t.Logf("length test; expected %v got %v", test.Output, actual)
		t.Fail()
		return
	}
	for i, itm := range actual {
		if test.Output[i] != itm {
			t.Logf("expected %d at index %d, got %d", test.Output[i], i, itm)
			t.Fail()
		}
	}
}
