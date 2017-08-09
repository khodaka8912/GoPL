package intset

import "testing"

func TestIntSet_Add(t *testing.T) {
	s := IntSet{}
	m := map[int]bool{}

	ints := []int{0, 1, 31, 32, 64}
	for _, i := range ints {
		s.Add(i)
		m[i] = true
	}

	for i := 0; i < 1024; i++ {
		if s.Has(i) != m[i] {
			t.Errorf("Has(%d) is diffrent from map", i)
		}
	}
}
