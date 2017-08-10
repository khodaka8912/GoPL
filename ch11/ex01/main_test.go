package main

import (
	"bytes"
	"testing"
)

func TestCharcount(t *testing.T) {
	tests := []struct {
		in      []byte
		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		{
			[]byte("Hello, World"),
			map[rune]int{'H': 1, 'e': 1, 'l': 3, 'o': 2, ',': 1, ' ': 1, 'W': 1, 'r': 1, 'd': 1},
			[]int{0, 12, 0, 0, 0},
			0,
		}, {
			[]byte("\300Hęllō, 世界。😃"),
			map[rune]int{'H': 1, 'ę': 1, 'l': 2, 'ō': 1, ',': 1, ' ': 1, '世': 1, '界': 1, '。': 1, '😃': 1},
			[]int{0, 5, 2, 3, 1},
			1,
		},
	}
	for _, test := range tests {
		counts, utflen, invalid, err := charcount(bytes.NewReader(test.in))
		if err != nil {
			t.Errorf("%v", err)
			continue
		}

		if len(counts) != len(test.counts) {
			t.Errorf("len(counts) is %d, want %d", len(counts), len(test.counts))
		}
		for k, v := range test.counts {
			count, ok := counts[k]
			if !ok {
				t.Errorf("counts does not contain %q", k)
			}
			if count != v {
				t.Errorf("counts[%q] = %d, want %d", k, count, v)
			}
		}

		if len(utflen) != len(test.utflen) {
			t.Errorf("len(utflen) = %d, want %d\n", len(utflen), len(test.utflen))
		}
		for i, count := range utflen {
			if count != test.utflen[i] {
				t.Errorf("utflen[%d] = %d, want %d\n", i, count, test.utflen[i])
			}
		}

		if invalid != test.invalid {
			t.Errorf("invalid = %d, want %d\n", invalid, test.invalid)
		}

	}
}
