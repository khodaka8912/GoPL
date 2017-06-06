package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("needs 2 args at least.")
	}
	fmt.Println(join(args[0], args[1:]...))
}

func join(sep string, values ...string) string {
	if len(values) == 0 {
		return ""
	}
	buf := bytes.NewBufferString(values[0])
	for _, s := range values[1:] {
		buf.WriteString(sep)
		buf.WriteString(s)
	}
	return buf.String()
}
