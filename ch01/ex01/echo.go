package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(getJoinedArgs(os.Args))
}

func getJoinedArgs(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}
