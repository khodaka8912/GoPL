package main

import (
	"bytes"
	"fmt"
	"os"
)

var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra": true},

	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":  {"discrete math": true},
	"databases":        {"data structures": true},
	"discrete math":    {"intro to programming": true},
	"formal languages": {"discrete math": true},
	"networks":         {"operating systems": true},
	"operating systems": {
		"data structures":       true,
		"computer organization": true,
	},
	"programming languages": {
		"data structures":       true,
		"computer organization": true,
	},
	"linear algebra": {"calculus": true},
}

func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sort failed:%v", err)
		os.Exit(1)
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) ([]string, error) {
	order := []string{}
	seen := map[string]bool{}
	var visitAll func(items map[string]bool, route []string) error
	visitAll = func(items map[string]bool, route []string) error {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				err := visitAll(m[item], append(route, item))
				if err != nil {
					return err
				}
				order = append(order, item)
			} else {
				for _, s := range route {
					if item == s {
						return fmt.Errorf("circuration detected :%v\n", sequenceString(append(route, item)))
					}
				}
			}
		}
		return nil
	}
	keys := map[string]bool{}
	for item := range m {
		keys[item] = true
	}
	err := visitAll(keys, []string{})
	return order, err
}

func sequenceString(s []string) string {
	if len(s) == 0 {
		return ""
	}
	buf := bytes.NewBufferString(fmt.Sprintf("[%q", s[0]))
	for _, str := range s[1:] {
		fmt.Fprintf(buf, " -> %q", str)
	}
	buf.WriteRune(']')
	return buf.String()
}
