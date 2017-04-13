package main

import (
	"fmt"
	"os"
	"strings"
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

func getJoinedArgs2(args []string) string {
	return strings.Join(args, " ")
}
