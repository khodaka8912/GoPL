package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("%s -> %s\n", arg, comma(arg))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	buf := bytes.Buffer{}
	i := (n-1)%3 + 1
	buf.WriteString(s[:i])
	for s = s[i:]; len(s) > 0; s = s[3:] {
		buf.WriteRune(',')
		buf.WriteString(s[:3])
	}
	return buf.String()
}
