package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		word     string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"abc", false},
		{"abcba", true},
		{"abccba", true},
	}
	for _, test := range tests {
		if actual := IsPalindrome(word(test.word)); actual != test.expected {
			t.Errorf("IsPalindrome(%s) = %b, want %b", test.word, actual, test.expected)
		}
	}
}
