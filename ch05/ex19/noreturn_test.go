package main

import "testing"

func TestReturnLessEcho(t *testing.T) {
	tests := []string{"", "abc", "Hello, 世界"}
	for _, test := range tests {
		if echo := returnLessEcho(test); echo != test {
			t.Errorf("returnLessEcho(%q) = %q, want %q", test, echo, test)
		}
	}
}
