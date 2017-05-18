package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode"
)

type categories struct {
	letter  int
	number  int
	mark    int
	space   int
	control int
	punct   int
	symbol  int
	other   int
}

func main() {
	s := "\u0300Hello\b, \t世界。2017😃"
	counts := charCount(bufio.NewReader(bytes.NewReader([]byte(s))))
	fmt.Printf("letters=%d\n", counts.letter)
	fmt.Printf("numbers=%d\n", counts.number)
	fmt.Printf("marks=%d\n", counts.mark)
	fmt.Printf("spaces=%d\n", counts.space)
	fmt.Printf("controls=%d\n", counts.control)
	fmt.Printf("puncts=%d\n", counts.punct)
	fmt.Printf("symbols=%d\n", counts.symbol)
	fmt.Printf("others=%d\n", counts.other)
}

func charCount(in *bufio.Reader) categories {
	counts := categories{}
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		switch {
		case unicode.IsSpace(r):
			counts.space++
		case unicode.IsPunct(r):
			counts.punct++
		case unicode.IsControl(r):
			counts.control++
		case unicode.IsMark(r):
			counts.mark++
		case unicode.IsSymbol(r):
			counts.symbol++
		case unicode.IsNumber(r):
			counts.number++
		case unicode.IsLetter(r):
			counts.letter++
		default:
			counts.other++
		}
	}
	return counts
}
