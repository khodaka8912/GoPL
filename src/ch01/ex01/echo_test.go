package main

import (
	"testing"
)

func TestGetJoinedArgs(t *testing.T) {
	cases := [][]string{{}, {"echo", "hello", "go"}}
	expected := []string{"", "echo hello go"}
	for n, arg := range cases {
		actual := getJoinedArgs(arg)
		if actual != expected[n] {
			t.Errorf("getJoinedArgs(%q) want %q but %q", arg, expected[n], actual)
		}
	}
}
