package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	var depthV2 int = 0
	var startElementV2 func(n *html.Node)
	var endElementV2 func(n *html.Node)

	startElementV2 = func(n *html.Node) {
		if n.Type == html.ElementNode {
			depthV2++
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		}
	}

	endElementV2 = func(n *html.Node) {
		if n.Type == html.ElementNode {
			depthV2--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	forEachNodeV2(doc, startElementV2, endElementV2)
}

func forEachNodeV2(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNodeV2(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
