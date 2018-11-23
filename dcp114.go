package daily_coding_practice

import (
	"strings"
)

// Run executes the algorithm and returns the result
func Run(input string, separators string) string {
	values := split(input, separators)

	if len(values) == 1 {
		return input
	}

	startBack := len(values) - 1

	for i, value := range values {
		//;fmt.Printf("value := %s\n", value)
		if "" != value && strings.Contains(separators, value) {
			continue
		}

		memorize := value
		for fromBack := startBack; fromBack > i; fromBack-- {
			startBack--
			if strings.Contains(separators, values[fromBack]) {
				continue
			}
			values[i] = values[fromBack]
			values[fromBack] = memorize
			break
		}
	}
	return strings.Join(values, "")
}

// split  breaks string into pieces giving each separator a spot
func split(input string, separators string) []string {
	var values []string
	value := ""
	lastOneSeparator := false

	for _, r := range input {
		character := string(r)
		if strings.Contains(separators, character) {
			value, values, lastOneSeparator = handleSeparator(value, values, lastOneSeparator, character)
		} else {
			value = value + string(r)
			lastOneSeparator = false
		}
	}
	if value != "" {
		values = append(values, value)
	}

	return values
}

func handleSeparator(value string, values []string, lastOneSeparator bool, character string) (string, []string, bool) {
	if len(value) > 0 {
		values = append(values, value)
	}
	if lastOneSeparator {
		values = append(values, "")
	}
	values = append(values, character)
	value = ""
	lastOneSeparator = true
	return value, values, lastOneSeparator
}
