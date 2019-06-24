package daily_coding_practice

func DCP331(input string) int {
	count := 0
	numberOfX := countInstancesOf('x', input)
	for pos, char := range input {
		if pos < numberOfX && char != 'x' {
			count = count + 1
		}
		if pos >= numberOfX && char != 'y' {
			count = count + 1
		}
	}
	return count
}

func countInstancesOf(of rune, input string) int {
	count := 0
	for _, char := range input {
		if char == of {
			count = count + 1
		}
	}
	return count
}
