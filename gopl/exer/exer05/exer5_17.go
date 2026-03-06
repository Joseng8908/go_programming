package main

import "golang.org/x/net/html"

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var result []*html.Node
	if doc == nil {
		return nil
	}
	if doc.Type == html.ElementNode {
		for _, n := range name {
			if doc.Data == n {
				result = append(result, doc)
				break
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, ElementsByTagName(c, name...)...)
	}
	return result
}
