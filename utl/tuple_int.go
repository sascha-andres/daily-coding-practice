package utl

type TupleInt struct {
	left  int
	right int
}

func (t *TupleInt) Left() int {
	return t.left
}

func (t *TupleInt) Right() int {
	return t.right
}

// NewTupleInt returns a new int tuple
func NewTupleInt(left, right int) *TupleInt {
	return &TupleInt{
		left:  left,
		right: right,
	}
}
