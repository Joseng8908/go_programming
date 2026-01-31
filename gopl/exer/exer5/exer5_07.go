package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth = 0

func Outline() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	forEachNodeV1(doc, startElement, endElement)
}
func forEachNodeV1(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNodeV1(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		var s strings.Builder
		s.WriteString("<")
		s.WriteString(n.Data)
		for _, attr := range n.Attr {
			s.WriteString(" ")
			s.WriteString(attr.Key)
			s.WriteString("=\"")
			s.WriteString(attr.Val)
			s.WriteString("\"")
		}

		if n.FirstChild == nil { // 자식이 없을 때
			fmt.Printf("%*s%s/>\n", depth*2, "", s.String())
		} else { // 자식이 있을 때
			fmt.Printf("%*s%s>\n", depth*2, "", s.String())
			depth++
		}
	} else if n.Type == html.TextNode {
		trimmedData := strings.TrimSpace(n.Data)
		if trimmedData != "" {
			// 여기도 n.Data 대신 trimmedData를 사용해야 공백이 제거돼요!
			fmt.Printf("%*s%s\n", depth*2, "", trimmedData)
		}
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	// ElementNode이고, 자식이 있었던 태그에 대해서만 닫아줍니다.
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

// exer5_8
func ElementById(doc *html.Node, id string) *html.Node {
	if doc.Type == html.ElementNode {
		for _, attr := range doc.Attr {
			if attr.Key == "id" && attr.Val == id {
				return doc
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		if e := ElementById(c, id); e != nil {
			return e
		}
	}
	return nil
}
