package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	info := make(map[string]int)
	createElementInfo(info, doc)
	fmt.Println("Element Info:")
	for name, count := range info {
		fmt.Printf("%s: %d\n", name, count)
	}
}
func createElementInfo(info map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		info[n.Data]++
	}
	createElementInfo(info, n.FirstChild)
	createElementInfo(info, n.NextSibling)
}
