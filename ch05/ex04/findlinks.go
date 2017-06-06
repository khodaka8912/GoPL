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
	for _, link := range visitLinks(nil, doc) {
		fmt.Println(link)
	}
}

var elements = map[string]string{
	"a":      "href",
	"img":    "src",
	"link":   "href",
	"script": "src",
}

func visitLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && elements[n.Data] != "" {
		for _, a := range n.Attr {
			if a.Key == elements[n.Data] {
				fmt.Printf("find <%s %s=%s>\n", n.Data, a.Key, a.Val)
				links = append(links, a.Val)
			}
		}
	}
	if c := n.FirstChild; c != nil {
		links = visitLinks(links, c)
	}
	if s := n.NextSibling; s != nil {
		links = visitLinks(links, s)
	}
	return links
}
