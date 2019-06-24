package utl

// LongestCommonSubsequence https://en.wikipedia.org/wiki/Longest_common_subsequence_problem#First_property
func LongestCommonSubsequence(first, second string, firstIndex, secondIndex int) string {
	if firstIndex == len(first) || secondIndex == len(second) {
		return ""
	}
	if first[firstIndex] == second[secondIndex] {
		return string(first[firstIndex]) + LongestCommonSubsequence(first, second, firstIndex+1, secondIndex+1)
	}
	resultA := LongestCommonSubsequence(first, second, firstIndex+1, secondIndex)
	resultB := LongestCommonSubsequence(first, second, firstIndex, secondIndex+1)
	if len(resultA) > len(resultB) {
		return resultA
	}
	return resultB
}
