package intset

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) addAll(values ...int) {
	for _, x := range values {
		s.Add(x)
	}
}

// 要素数を返します
func (s *IntSet) Len() (len int) {
	for _, word := range s.words {
		len += popCount(word)
	}
	return len
}

// セットからxを取り除きます
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

// セットからすべての要素を取り除きます
func (s *IntSet) Clear() {
	s.words = []uint64{}
}

// セットからすべての要素を取り除きます
func (s *IntSet) Copy() *IntSet {
	copy := &IntSet{}
	copy.UnionWith(s)
	return copy
}

func NewIntSet(values ...int) *IntSet {
	s := &IntSet{}
	for _, val := range values {
		s.Add(val)
	}
	return s
}

func popCount(x uint64) int {
	const (
		m1  = 0x5555555555555555
		m2  = 0x3333333333333333
		m4  = 0x0f0f0f0f0f0f0f0f
		m8  = 0x00ff00ff00ff00ff
		m16 = 0x0000ffff0000ffff
		m32 = 0x00000000ffffffff
	)
	x = (x & m1) + ((x >> 1) & m1)
	x = (x & m2) + ((x >> 2) & m2)
	x = (x & m4) + ((x >> 4) & m4)
	x = (x & m8) + ((x >> 8) & m8)
	x = (x & m16) + ((x >> 16) & m16)
	x = (x & m32) + ((x >> 32) & m32)
	return int(x)
}
