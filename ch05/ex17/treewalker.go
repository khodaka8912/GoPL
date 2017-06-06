package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("needs args [url] [tag name]...")
		return
	}
	doc, err := fetchAndParse(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	elems := ElementByTagName(doc, args[1:]...)
	fmt.Printf("%d tags are found.\n", len(elems))
	for _, n := range elems {
		fmt.Println(tagString(n))
	}
}

func fetchAndParse(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func ElementByTagName(doc *html.Node, names ...string) []*html.Node {
	found := []*html.Node{}
	startElement := func(n *html.Node) bool {
		if n.Type == html.ElementNode && contains(names, n.Data) {
			found = append(found, n)
		}
		return true
	}
	forEachNode(doc, startElement, nil)
	return found
}

func contains(list []string, str string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
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

func tagString(n *html.Node) string {
	buf := &bytes.Buffer{}
	buf.WriteRune('<')
	buf.WriteString(n.Data)
	for _, a := range n.Attr {
		fmt.Fprintf(buf, " %s=%q", a.Key, a.Val)
	}
	buf.WriteRune('>')
	return buf.String()
}
