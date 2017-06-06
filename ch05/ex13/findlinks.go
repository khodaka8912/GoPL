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

	"gopl.io/ch5/links"
)

var cd string

func main() {
	args := os.Args[1:]
	dir, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	cd = dir

	if len(args) == 0 {
		return
	}
	breadthFirst(crawl, args)
}

var elements = map[string]string{
	"a":      "href",
	"img":    "src",
	"link":   "href",
	"script": "src",
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
