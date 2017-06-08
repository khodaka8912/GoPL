package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	word := word(os.Args[1:])
	if IsPalindrome(word) {
		fmt.Printf("%v is palindrome.\n", word)
	} else {
		fmt.Printf("%v is not palindrome.\n", word)
	}
}

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

type word []string

func (w word) Len() int {
	return len(w)
}

func (w word) Less(i, j int) bool {
	return w[i] < w[j]
}

func (w word) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
