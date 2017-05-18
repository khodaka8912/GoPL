package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

var n = flag.Int("n", 0, "count of rotation")

func main() {
	flag.Parse()
	args := flag.Args()
	var s []int
	for _, arg := range args {
		i, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatal(err)
			return
		}
		s = append(s, i)
	}
	rotate(s, *n)
	fmt.Println(s)
}

func rotate(s []int, n int) {
	size := len(s)
	if n > size {
		n %= size
	}
	reverse(s)
	reverse(s[:n])
	reverse(s[n:])
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
