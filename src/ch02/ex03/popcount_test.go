package popcount

import (
	"testing"
)

var values = []uint64{0x0, 0xffffffffffffffff, 0x5555555555555555}
var counts = []int{0, 64, 32}

func TestPopCount(t *testing.T) {
	for i, value := range values {
		result := PopCount(value)
		if result != counts[i] {
			t.Errorf("PopCount(%d) want %d but %d", value, counts[i], result)
		}
	}
}

func TestPopCount2(t *testing.T) {
	for i, value := range values {
		result := PopCount2(value)
		if result != counts[i] {
			t.Errorf("PopCount(%d) want %d but %d", value, counts[i], result)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(values[i%3])
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(values[i%3])
	}
}
