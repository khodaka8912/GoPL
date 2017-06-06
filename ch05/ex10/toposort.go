package main

import "fmt"

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
}

func main() {
	order := topoSort(prereqs)
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) []string {
	order := []string{}
	seen := map[string]bool{}
	var visitAll func(items map[string]bool)
	visitAll = func(items map[string]bool) {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	keys := map[string]bool{}
	for item := range m {
		keys[item] = true
	}
	visitAll(keys)
	return order
}
