package daily_coding_practice

import "errors"

func Run165(input []int) ([]int, error) {
	if len(input) == 0 {
		return nil, errors.New("empty array passed")
	}

	result := make([]int, len(input))

	for i, itm := range input {
		cnt, err := countSmaller165(input, itm, i+1)
		if err != nil {
			return nil, err
		}
		result[i] = cnt
	}

	return result, nil
}

func countSmaller165(input []int, number, index int) (int, error) {
	if index > len(input) {
		return 0, errors.New("index out of bounds")
	}
	if index == len(input) {
		return 0, nil
	}
	cnt := 0
	for idx := index; idx < len(input); idx++ {
		if input[idx] < number {
			cnt++
		}
	}
	return cnt, nil
}
