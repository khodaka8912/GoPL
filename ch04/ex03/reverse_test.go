package main

import "testing"

func TestReverse(t *testing.T) {
	in, out := [size]int{1, 2, 3, 4, 5}, [size]int{5, 4, 3, 2, 1}
	reverse(&in)
	for i := 0; i < size; i++ {
		if in[i] != out[i] {
			t.Errorf("reversed in[%d] wants %d but %d\n", i, in[i], out[i])
		}
	}

}
