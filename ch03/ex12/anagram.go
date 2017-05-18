package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		an := isAnagram(args[0], args[1])
		fmt.Printf("isAnagram:%t\n", an)
	} else {
		fmt.Println("usage anagram str1 str2")
	}
}

func isAnagram(s1, s2 string) bool {
	runes1, runes2 := countByRunes(s1), countByRunes(s2)
	if len(runes1) != len(runes2) {
		return false
	}
	for k, v := range runes1 {
		if runes2[k] != v {
			return false
		}
	}
	return true
}

func countByRunes(s string) map[rune]int {
	runes := make(map[rune]int)
	for _, r := range s {
		runes[r]++
	}
	return runes
}
