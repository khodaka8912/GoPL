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

func TestGetJoinedArgs2(t *testing.T) {
	cases := [][]string{{}, {"echo", "hello", "go"}}
	expected := []string{"", "echo hello go"}
	for n, arg := range cases {
		actual := getJoinedArgs2(arg)
		if actual != expected[n] {
			t.Errorf("getJoinedArgs(%q) want %q but %q", arg, expected[n], actual)
		}
	}
}

var values = [][]string{
	{"a", "b", "c", "d","e","f","g","h"},
	{"123456789", "123456789", "123456789", "123456789", "123456789", "123456789", "123456789",
		"123456789", "123456789", "123456789", "123456789", "123456789", "123456789", "123456789", },
	{"1", "2", "3"},
}

func BenchmarkGetJoinedArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getJoinedArgs(values[i%3])
	}

}

func BenchmarkGetJoinedArgs2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getJoinedArgs2(values[i%3])
	}

}
