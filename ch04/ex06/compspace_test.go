package main

import "testing"

func TestReverse(t *testing.T) {
	in, out := []byte("a b\tc　\t def    g"), []byte("a b c def g")
	actual := compressSpace(in)
	if len(actual) != len(out) {
		t.Errorf("compress(%q) wants %q but %q", in, out, actual)
	}
	for i := 0; i < len(out); i++ {
		if actual[i] != out[i] {
			t.Errorf("compress(%q)[%d] wants %q but %q\n", string(in), i, out[i], actual[i])
		}
	}

}
