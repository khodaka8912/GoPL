package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for k, v := range countElements(make(map[string]int), doc) {
		fmt.Printf("<%s> : %d\n", k, v)
	}
}

func countElements(elements map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}
	if c := n.FirstChild; c != nil {
		countElements(elements, c)
	}
	if s := n.NextSibling; s != nil {
		countElements(elements, s)
	}
	return elements
}
