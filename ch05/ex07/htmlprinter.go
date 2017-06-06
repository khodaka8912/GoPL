package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("needs url as arguments.")
		return
	}
	for _, url := range args {
		err := PrintHtml(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error counting %q :%s", url, err)
			continue
		}
	}
}

func PrintHtml(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	printHtml(doc, os.Stdout)
	return nil
}

func printHtml(doc *html.Node, out io.Writer) {
	startElement := func(n *html.Node) {
		switch n.Type {
		case html.ElementNode:
			fmt.Fprintf(out, "<%s", n.Data)
			for _, a := range n.Attr {
				fmt.Fprintf(out, " %s=\"%s\"", a.Key, a.Val)
			}
			if n.FirstChild == nil {
				fmt.Fprintf(out, "/>")
			} else {
				fmt.Fprintf(out, ">")
			}
		case html.TextNode:
			fmt.Fprintf(out, n.Data)
		case html.CommentNode:
			fmt.Fprintf(out, "<!-- %s -->", n.Data)
		}
	}
	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode && n.FirstChild != nil {
			fmt.Fprintf(out, "</%s>", n.Data)
		}
	}
	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
