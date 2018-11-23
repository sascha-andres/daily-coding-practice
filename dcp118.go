package daily_coding_practice

import "sort"

// Given a sorted list of integers, square the elements and give the output in sorted order.
func Run118(input []int) []int {
	result := make([]int, len(input))
	for i, val := range input {
		result[i] = val * val
	}
	sort.Ints(result)
	return result
}
