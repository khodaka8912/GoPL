package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	matchers := parseMatchers(os.Args[1:])
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, matchers) {
				fmt.Printf("%s\n", tok)
			}
		}
	}
}
func parseMatchers(args []string) []tagMatcher {
	matchers := []tagMatcher{}
	for _, arg := range args {
		n := strings.Index(arg, "=")
		if n > 0 && n < len(arg)-1 {
			matcher := attr{arg[:n], arg[n+1:]}
			matchers = append(matchers, matcher)
		} else {
			matchers = append(matchers, name(arg))
		}
	}
	return matchers
}

type tagMatcher interface {
	match(element xml.StartElement) bool
}

type name string

func (n name) match(element xml.StartElement) bool {
	return name(element.Name.Local) == n
}

type attr struct {
	name  string
	value string
}

func (a attr) match(element xml.StartElement) bool {
	for _, attr := range element.Attr {
		if attr.Name.Local == a.name && attr.Value == a.value {
			return true
		}
	}
	return false
}

func containsAll(x []xml.StartElement, y []tagMatcher) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if y[0].match(x[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
