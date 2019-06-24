package utl

import "testing"

var testCasesLCS = [][]string{
	{"ABCD", "ACDEB", "ACD"},
	{"I've been waiting", "I go waiting", "I  waiting"},
}

func TestLongestCommonSubsequence(t *testing.T) {
	for _, val := range testCasesLCS {
		result := LongestCommonSubsequence(val[0], val[1], 0, 0)
		if result != val[2] {
			t.Logf("Expected [%s], got [%s] for [%s] and [%s]", val[2], result, val[0], val[1])
			t.Fail()
		}
	}
}
