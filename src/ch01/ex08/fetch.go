package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	result := fetch(os.Args[1:], os.Stdout)
	if result != 0 {
		os.Exit(result)
	}
}

// return 0 if all succeeded, or 1 if an error has occurred
func fetch(urls []string, writer io.Writer) int {
	const proto = "http://"
	for _, url := range urls {
		if !strings.HasPrefix(url, proto) {
			url = proto + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			return 1
		}
		_, err = io.Copy(writer, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			return 1
		}
	}
	return 0
}
