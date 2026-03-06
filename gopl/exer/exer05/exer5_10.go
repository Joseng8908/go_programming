package main

import (
	"fmt"
	"os"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func exer5_10() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d: \t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visiting := make(map[string]bool)

	visitAll = func(items []string) {
		for _, item := range items {
			if visiting[item] {
				fmt.Printf("cycle: %s\n", item)
				os.Exit(1)
			}
			visiting[item] = true
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
			visiting[items[0]] = false
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	visitAll(keys)
	return order
}
