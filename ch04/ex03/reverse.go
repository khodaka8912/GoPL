package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const size = 5

func main() {
	args := os.Args[1:]
	if len(args) != size {
		fmt.Printf("reverse([%d]int) needs just %[1]d args\n", size)
		return
	}
	s := [size]int{}
	for n, arg := range args {
		i, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatal(err)
			return
		}
		s[n] = i
	}
	reverse(&s)
	fmt.Println(s)
}

func reverse(s *[size]int) {
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
