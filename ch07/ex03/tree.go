package main

import (
	"bytes"
	"fmt"
)

func main() {
	t := &tree{
		10,
		&tree{
			5,
			&tree{
				3,
				&tree{
					1,
					nil,
					nil,
				},
				&tree{
					4,
					nil,
					nil,
				},
			},
			&tree{
				8,
				nil,
				&tree{
					9,
					nil,
					nil,
				},
			},
		},
		&tree{
			20,
			&tree{
				15,
				&tree{
					12,
					nil,
					&tree{
						13,
						nil,
						nil,
					},
				},
				nil,
			},
			nil,
		},
	}
	fmt.Println(t)
	fmt.Println(t.String2())
}

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	return fmt.Sprintf("%v", appendValues([]int{}, t))
}

// without making slice
func (t *tree) String2() (s string) {
	buf := &bytes.Buffer{}
	buf.WriteRune('[')
	t.writeInnerString(buf)
	buf.WriteRune(']')
	return buf.String()
}

func (t *tree) writeInnerString(b *bytes.Buffer) {
	if t.left != nil {
		t.left.writeInnerString(b)
	}
	if b.Len() > 1 {
		b.WriteRune(' ')
	}
	fmt.Fprintf(b, "%d", t.value)
	if t.right != nil {
		t.right.writeInnerString(b)
	}
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
