package popcount

import (
	"testing"
)

var values = []uint64{0x0, 0x5555555555555555, 0xffffffffffffffff}
var counts = []int{0, 32, 64}

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

func benchmarkPopCount(b *testing.B, value uint64) {
	for i := 0; i < b.N; i++ {
		PopCount(value)
	}
}

func benchmarkPopCount2(b *testing.B, value uint64) {
	for i := 0; i < b.N; i++ {
		PopCount2(value)
	}
}

func benchmarkPopCount3(b *testing.B, value uint64) {
	for i := 0; i < b.N; i++ {
		PopCount3(value)
	}
}

func benchmarkPopCount4(b *testing.B, value uint64) {
	for i := 0; i < b.N; i++ {
		PopCount4(value)
	}
}

func BenchmarkPopCount_0(b *testing.B) {
	benchmarkPopCount(b, values[0])
}

func BenchmarkPopCount2_0(b *testing.B) {
	benchmarkPopCount2(b, values[0])
}

func BenchmarkPopCount3_0(b *testing.B) {
	benchmarkPopCount3(b, values[0])
}

func BenchmarkPopCount4_0(b *testing.B) {
	benchmarkPopCount4(b, values[0])
}

func BenchmarkPopCount_32(b *testing.B) {
	benchmarkPopCount(b, values[1])
}

func BenchmarkPopCount2_32(b *testing.B) {
	benchmarkPopCount2(b, values[1])
}

func BenchmarkPopCount3_32(b *testing.B) {
	benchmarkPopCount3(b, values[1])
}

func BenchmarkPopCount4_32(b *testing.B) {
	benchmarkPopCount4(b, values[1])
}

func BenchmarkPopCount_64(b *testing.B) {
	benchmarkPopCount(b, values[2])
}

func BenchmarkPopCount2_64(b *testing.B) {
	benchmarkPopCount2(b, values[2])
}

func BenchmarkPopCount3_64(b *testing.B) {
	benchmarkPopCount3(b, values[2])
}

func BenchmarkPopCount4_64(b *testing.B) {
	benchmarkPopCount4(b, values[2])
}
