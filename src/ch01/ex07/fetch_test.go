package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetch(t *testing.T) {
	const response = "Test Response"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, response)
	}))
	out := new(bytes.Buffer)
	arg := []string{server.URL}
	result := fetch(arg, out)
	if result != 0 {
		t.Errorf("fetch(%q) reuslt want %q but %q", arg, 0, result)
	}
	content := out.String()
	if content != response {
		t.Errorf("fetch(%q) writer want %q but %q", arg, response, content)
	}
	server.Close()
}
