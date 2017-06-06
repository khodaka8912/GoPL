package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("needs url as arguments.")
		return
	}
	for _, url := range args {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error counting %q :%s", url, err)
			continue
		}
		fmt.Printf("counts of %q\n", url)
		fmt.Printf("words=%d, images=%d\n", words, images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}
func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	if n.Type == html.TextNode {
		input := bufio.NewScanner(strings.NewReader(n.Data))
		input.Split(bufio.ScanWords)
		for input.Scan() {
			input.Text()
			words++
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cwords, cimages := countWordsAndImages(c)
		words, images = words+cwords, images+cimages
	}
	return words, images
}
