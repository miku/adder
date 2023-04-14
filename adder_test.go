package adder

import "testing"

func TestAdd(t *testing.T) {
	// Table-driven tests
	var cases = []struct {
		a      int
		b      int
		result int
	}{
		{a: 1, b: 2, result: 3},
		{a: 2, b: 2, result: 4},
		{a: 0, b: 1, result: 1},
	}
	t.Logf("testing %v cases", len(cases))
	for _, c := range cases {
		result := Add(c.a, c.b)
		if result != c.result {
			t.Fatalf("got %v, want %v", result, c.result)
		}
	}
}
