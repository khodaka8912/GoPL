package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("needs args [srcString] [replaceString]\n replaceString can use '$1'")
		return
	}
	replaced := expand(args[0], func(s string) string {
		return strings.Replace(args[1], "$1", s, -1)
	})
	fmt.Printf("%s -> %s", args[0], replaced)
}

func expand(s string, f func(string) string) string {
	token := regexp.MustCompile(`\$[a-zA-Z_]+`)
	return token.ReplaceAllStringFunc(s, f)
}
