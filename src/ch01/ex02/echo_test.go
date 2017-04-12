package main

import (
	"testing"
)

func TestGetJoinedArgs(t *testing.T) {
	cases := [][]string{{}, {"echo", "hello", "go"}}
	expected := []string{"", "0: echo\n1: hello\n2: go"}
	for n, arg := range cases {
		actual := toLabeledLines(arg)
		if actual != expected[n] {
			t.Errorf("GetJoinedArgs(%q) want %q but %q", arg, expected[n], actual)
		}
	}
}
