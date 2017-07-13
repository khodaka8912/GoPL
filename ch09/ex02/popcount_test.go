package popcount

import (
	"testing"
)

var tests = []struct {
	value uint64
	count int
}{
	{0x0, 0},
	{0xffffffffffffffff, 64},
	{0x5555555555555555, 32},
	{0x8000000000000000, 1},
}

func TestPopCount(t *testing.T) {
	for _, test := range tests {
		result := PopCount(test.value)
		if result != test.count {
			t.Errorf("PopCount(%d) want %d but %d", test.value, test.count, result)
		}
	}
}
