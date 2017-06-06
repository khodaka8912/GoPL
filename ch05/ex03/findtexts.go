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
	for _, link := range visitTexts(nil, doc) {
		fmt.Println(link)
	}
}

func visitTexts(texts []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		if len(n.Data) > 0 {
			texts = append(texts, n.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "style" || c.Data == "script" {
			continue
		}
		texts = visitTexts(texts, c)
	}
	return texts
}
