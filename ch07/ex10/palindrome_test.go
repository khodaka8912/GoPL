package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		word     []string
		expected bool
	}{
		{[]string{}, true},
		{[]string{"a"}, true},
		{[]string{"a", "b", "c"}, false},
		{[]string{"a", "b", "c", "b", "a"}, true},
		{[]string{"a", "b", "c", "c", "b", "a"}, true},
	}
	for _, test := range tests {
		if actual := IsPalindrome(word(test.word)); actual != test.expected {
			t.Errorf("IsPalindrome(%s) = %b, want %b", test.word, actual, test.expected)
		}
	}
}
