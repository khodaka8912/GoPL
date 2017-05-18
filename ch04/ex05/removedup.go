package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(removeDup(os.Args[1:]))
}

func removeDup(s []string) []string {
	if len(s) < 2 {
		return s
	}
	n := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[n] {
			n++
			s[n] = s[i]
		}
	}
	return s[:n+1]
}
