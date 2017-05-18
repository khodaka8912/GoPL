package main

import "testing"

func TestReverse(t *testing.T) {
	in, out := []string{"a", "b", "b", "c", "x", "x", "x", "a"}, []string{"a", "b", "c", "x", "a"}
	actual := removeDup(in)
	if len(actual) != len(out) {
		t.Errorf("removeDup(%q) wants %q but %q", in, out, actual)
	}
	for i := 0; i < len(out); i++ {
		if actual[i] != out[i] {
			t.Errorf("removeDup(%q)[%d] wants %q but %q\n", in, i, out[i], actual[i])
		}
	}

}
