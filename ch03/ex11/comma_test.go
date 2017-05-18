package main

import "testing"

var testCases = []struct{ in, out string }{
	{"1", "1"},
	{"12", "12"},
	{"321", "321"},
	{"123456", "123,456"},
	{"66666666", "66,666,666"},
	{"7777777777777", "7,777,777,777,777"},
	{"-111", "-111"},
	{"12.3456", "12.345,6"},
	{"+7777777.777", "+7,777,777.777"},
}

func TestComma(t *testing.T) {
	for _, test := range testCases {

		if actual := comma(test.in); actual != test.out {
			t.Errorf("comma(%q) wants %q but %q\n", test.in, test.out, actual)
		}
	}
}
