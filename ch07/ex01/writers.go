package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	words := strings.Join(os.Args[1:], " ")
	fmt.Println(words)
	var w WordCounter
	index := 0
	for index < len(words) {
		size, err := w.Write([]byte(words[index:]))
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		index += size
	}
	fmt.Printf("%d words are written\n", w)
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	size, token, err := bufio.ScanWords(p, true)
	if token != nil {
		*c++
	}
	return size, err
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	size := 0
	for r, s := utf8.DecodeRune(p); len(p) > 0; p = p[s:] {
		if r == '\n' {
			*c++
		}
		size += s
	}
	return size, nil
}
