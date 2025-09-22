package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	//5.1
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}

func outline(stack []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
	return stack
}
