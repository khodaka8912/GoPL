package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts, size := wordfreq(os.Stdin)
	for w, n := range counts {
		percent := float64(n) / float64(size) * 100
		fmt.Printf("%s=%f%%\n", w, percent)
	}
}

func wordfreq(in *os.File) (map[string]int, int) {
	input := bufio.NewScanner(in)
	counts := map[string]int{}
	size := 0
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		counts[word]++
		size++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	return counts, size
}
