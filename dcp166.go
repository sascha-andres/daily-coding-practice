package daily_coding_practice

type Result166 struct {
	inputArrays [][]int

	outerArrayIndex int
	innerArrayIndex int
}

func NewResult166(input [][]int) *Result166 {
	return &Result166{
		inputArrays:     input,
		outerArrayIndex: -1,
		innerArrayIndex: -1,
	}
}

func (r *Result166) Next() int {
	return 0
}

func (r *Result166) HasNext() bool {
	if len(r.inputArrays) <= r.outerArrayIndex {
		return false
	}

	if len(r.inputArrays[r.outerArrayIndex+1]) >= r.innerArrayIndex {
		// if
	}

	return false
}

func Run166(input [][]int) *Result166 {
	return NewResult166(input)
}
