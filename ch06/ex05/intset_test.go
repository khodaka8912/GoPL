package intset

import (
	"testing"
)

func TestIntSet_Len(t *testing.T) {
	var tests = []struct {
		val *IntSet
		len int
	}{
		{NewIntSet(), 0},
		{NewIntSet(1), 1},
		{NewIntSet(0, 63, 64, 127, 128, 191), 6},
		{NewIntSet(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
			17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
			33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63), 64},
	}

	for _, test := range tests {
		if actual := test.val.Len(); actual != test.len {
			t.Errorf("IntSet(%v).Len() wants %d but %d", test.val, test.len, actual)
		}
	}
}

func TestIntSet_Remove(t *testing.T) {
	var tests = []struct {
		val    *IntSet
		x      int
		remain []int
	}{
		{NewIntSet(), 0, []int{}},
		{NewIntSet(1), 1, []int{}},
		{NewIntSet(1, 2, 3), 4, []int{1, 2, 3}},
		{NewIntSet(0, 63, 64, 127, 128, 191), 63, []int{0, 64, 127, 128, 191}},
	}
	for _, test := range tests {
		target := &IntSet{test.val.words}
		target.Remove(test.x)
		if target.Len() != len(test.remain) {
			t.Errorf("IntSet%v.Remove(%d) got Intset%v, want %v", test.val, test.x, target, test.remain)
		}
		for _, i := range test.remain {
			if !target.Has(i) {
				t.Errorf("IntSet%v.Remove(%d) got Intset%v, want %v", test.val, test.x, target, test.remain)
			}
		}
	}
}

func TestIntSet_Clear(t *testing.T) {
	target := NewIntSet(0, 63, 64, 127, 128, 191)
	if target.Clear(); target.Len() != 0 {
		t.Errorf("Intset.Clear() got %v want {}", target)
	}
}

func TestIntSet_Copy(t *testing.T) {
	target := NewIntSet(0, 63, 64, 127, 128, 191)
	copied := target.Copy()
	if &target.words == &copied.words {
		t.Error("IntSet.Copy() didn't copy inner value. words points same address.")
	}
	if target.Len() != copied.Len() {
		t.Errorf("IntSet%v.Copy() got Intset%v, want Intset%v", target, copied, target)
	}
	for i, w := range copied.words {
		if target.words[i] != w {
			t.Errorf("IntSet%v.Copy() got Intset%v, want Intset%v", target, copied, target)
		}
	}

}
