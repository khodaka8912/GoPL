package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(toLabeledLines(os.Args))
}

func toLabeledLines(args []string) string {
	s, sep := "", ""
	for n, arg := range args {
		s += sep + strconv.Itoa(n) + ": " + arg
		sep = "\n"
	}
	return s
}
