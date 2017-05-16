package popcount

import (
	"testing"
)

var values = []uint64{0x0, 0xffffffffffffffff, 0x5555555555555555, 0x8000000000000000}
var counts = []int{0, 64, 32, 1}

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

func TestPopCount3(t *testing.T) {
	for i, value := range values {
		result := PopCount3(value)
		if result != counts[i] {
			t.Errorf("PopCount(%d) want %d but %d", value, counts[i], result)
		}
	}
}

func TestPopCount4(t *testing.T) {
	for i, value := range values {
		result := PopCount4(value)
		if result != counts[i] {
			t.Errorf("PopCount(%d) want %d but %d", value, counts[i], result)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(values[i%4])
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(values[i%4])
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(values[i%4])
	}
}

func BenchmarkPopCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount4(values[i%4])
	}
}
