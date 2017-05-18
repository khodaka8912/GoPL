package main

import "testing"

func TestReverse(t *testing.T) {
	in, out := []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}
	rotate(in, 3)
	for i := 0; i < len(in); i++ {
		if in[i] != out[i] {
			t.Errorf("rotated in[%d] wants %d but %d\n", i, in[i], out[i])
		}
	}

}
