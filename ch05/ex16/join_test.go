package main

import (
	"testing"
)

var tests = []struct {
	sep    string
	values []string
	out    string
}{
	{"abc", []string{""}, ""},
	{"abc", []string{"def"}, "def"},
	{"", []string{"a", "b", "c"}, "abc"},
	{"::", []string{"a", "b", "c"}, "a::b::c"},
}

func TestJoin(t *testing.T) {
	for _, test := range tests {
		if actual := join(test.sep, test.values...); actual != test.out {
			t.Errorf("join(%q, %q) wants %q but %q", test.sep, test.values, test.out, actual)
		}
	}
}
