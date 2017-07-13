package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/net/html"
	"gopl.io/ch5/links"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	breadthFirst(crawl, args)
}

func copy(filePath string, reader io.ReadCloser) error {
	defer reader.Close()
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, data, os.ModePerm)
	return err
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(urlStr string) []string {
	fmt.Println(urlStr)
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		log.Print(err)
	}
	domain := parsedUrl.Host
	resource := parsedUrl.Path
	if resource == "" || resource == "/" {
		resource = "/index.html"
	}
	filePath := filepath.FromSlash(domain + resource)
	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		log.Print(err)
	}
	resp, err := http.Get(urlStr)
	if err != nil {
		log.Print(err)
	} else {
		copy(filePath, resp.Body)
	}

	list, err := links.Extract(urlStr)
	if err != nil {
		log.Print(err)
	}
	return list
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
