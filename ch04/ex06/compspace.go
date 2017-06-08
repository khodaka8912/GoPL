package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	b := []byte("ab c   d \t e  _f g")
	b = compressSpace(b)
	fmt.Println(string(b))
}

func compressSpace(b []byte) []byte {
	s := string(b)
	var i int
	space := false
	for n, r := range s {
		if unicode.IsSpace(r) {
			if space {
				continue
			}
			b[i] = ' '
			i++
			space = true
		} else {
			size := utf8.RuneLen(r)
			copy(b[i:], b[n:n+size])
			i += size
			space = false
		}
	}
	return b[:i]
}
