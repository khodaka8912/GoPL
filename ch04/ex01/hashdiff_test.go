package main

import "testing"

var testCases = []struct {
	b1, b2 [32]byte
	out    int
}{
	{[32]byte{}, [32]byte{}, 0},
	{[32]byte{1, 2, 3, 4, 5}, [32]byte{}, 7},
	{[32]byte{1, 2, 3, 4, 5}, [32]byte{5, 4, 3, 2, 1}, 6},
}

func TestDiff(t *testing.T) {
	for _, test := range testCases {
		if actual := diff(&test.b1, &test.b2); actual != test.out {
			t.Errorf("diff(%q, %q) wants %d but %d", test.b1, test.b2, test.out, actual)
		}
	}
}
