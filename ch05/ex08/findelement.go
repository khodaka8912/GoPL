package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("needs args [url] [elementId]")
		return
	}
	url, id := args[0], args[1]
	n, err := fetchAndElementByID(url, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error counting %q :%s", url, err)
		os.Exit(1)
	}
	if n == nil {
		fmt.Printf("node of id:%q not found", id)
	}
	fmt.Printf("found node:<%s id=\"%s\">", n.Data, id)
}

func fetchAndElementByID(url string, id string) (found *html.Node, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return ElementByID(doc, id), nil
}

func ElementByID(doc *html.Node, id string) (found *html.Node) {
	startElement := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				found = n
				return false
			}
		}
		return true
	}
	forEachNode(doc, startElement, nil)
	return found
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil && !pre(n) {
		return false
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}

	if post != nil && !post(n) {
		return false
	}
	return true
}
