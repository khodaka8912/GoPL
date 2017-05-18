package main

import "testing"

var testCases = []struct {
	s1, s2 string
	out    bool
}{
	{"aaa", "aaa", true},
	{"aaa", "bbb", false},
	{"abcde", "aebdc", true},
	{"xyz", "xyzw", false},
	{"tom marvolo riddle", "iam lord voldemort", true},
}

func TestAnagram(t *testing.T) {
	for _, test := range testCases {
		if actual := isAnagram(test.s1, test.s2); actual != test.out {
			t.Errorf("isAnagram(%q, %q) wants %q but %q", test.s1, test.s2, test.out, actual)
		}
	}
}
