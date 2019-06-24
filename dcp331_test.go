package daily_coding_practice

import "testing"

var testsDCP331 = map[string]int{
	"xy":          0,
	"xx":          0,
	"yy":          0,
	"yx":          2,
	"xyxxxyxyy":   2,
	"xxyxxxyxxyy": 4,
}

var testsCountDCP331 = map[string]int{
	"xy": 1,
	"xx": 2,
	"yy": 0,
	"yx": 1,
}

func TestDCP133(t *testing.T) {
	for k, v := range testsDCP331 {
		t.Run(k, func(t2 *testing.T) {
			result := DCP331(k)
			if result != v {
				t2.Logf("expected [%d], got [%d]", v, result)
				t2.FailNow()
			}
		})
	}
}

func TestDCP331Count(t *testing.T) {
	for k, v := range testsCountDCP331 {
		t.Run(k, func(t2 *testing.T) {
			result := countInstancesOf('x', k)
			if result != v {
				t2.Logf("expected [%d], got [%d]", v, result)
				t2.FailNow()
			}
		})
	}
}
