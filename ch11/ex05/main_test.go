package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	for _, test := range []struct {
		s, sep string
		want   []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"a:b:c", ",", []string{"a:b:c"}},
		{"", ",", []string{""}},
		{"aaa/bb/c", "/", []string{"aaa", "bb", "c"}},
		{"x,y,z", "", []string{"x", ",", "y", ",", "z"}},
	} {
		got := strings.Split(test.s, test.sep)
		if len(got) != len(test.want) {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, len(got), len(test.want))
			continue
		}
		for i, word := range got {
			if word != test.want[i] {
				t.Errorf("Split(%q, %q)[%d] = %q, want %q", test.s, test.sep, i, word, test.want[i])
			}
		}

	}
}
