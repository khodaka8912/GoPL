package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

func TestFetch(t *testing.T) {
	const response = "Test Response"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, response)
	}))
	ch := make(chan string)
	url := server.URL
	go fetch(url, ch, 0)
	result := <-ch
	timeRegex, _ := regexp.Compile(`^[0-9\.]+s`)
	nbytes := fmt.Sprintf("%7d", len(response))
	if !timeRegex.MatchString(result) {
		t.Errorf("fetch(%q, ch) <-ch want to start with %q but %q", url, timeRegex, result)
	}
	if !strings.Contains(result, nbytes) {
		t.Errorf("fetch(%q, ch) <-ch want to contain %q but %q", url, nbytes, result)
	}
	if !strings.HasSuffix(result, url) {
		t.Errorf("fetch(%q, ch) <-ch want to end with %q but %q", url, url, result)
	}
	server.Close()
}
