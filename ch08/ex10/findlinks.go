package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func crawl(url string, ctx context.Context) []string {
	fmt.Println(url)
	list, err := Extract(url, ctx)
	if err != nil {
		log.Print(err)
	}
	return list
}
func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- os.Args[1:] }()
	cancel := make(chan struct{})
	ctx, cancelFunc := context.WithCancel(context.Background())
	n := 1
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				select {
				case <-cancel:
					return
				default:
					foundLinks := crawl(link, ctx)
					n++
					go func() { worklist <- foundLinks }()
				}
			}
		}()
	}

	input := bufio.NewScanner(os.Stdin)
	go func() {
		input.Scan()
		cancelFunc()
		close(cancel)
		fmt.Println("canceled")
	}()
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		select {
		case list := <-worklist:
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					unseenLinks <- link
				}
			}
		case <-cancel:
			return
		}
	}
}

func Extract(url string, ctx context.Context) ([]string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
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
