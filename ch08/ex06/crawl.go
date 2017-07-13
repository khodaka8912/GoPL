package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

type linkInfo struct {
	links []string
	depth int
}

var depth = flag.Int("depth", 0, "max depth to search links. 0 is unlimited.")

func main() {
	flag.Parse()
	worklist := make(chan linkInfo)
	var n int

	n++
	go func() { worklist <- linkInfo{flag.Args(), 0} }()
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		info := <-worklist
		if *depth > 0 && info.depth >= *depth {
			continue
		}
		for _, link := range info.links {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- linkInfo{crawl(link), info.depth + 1}
				}(link)
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
