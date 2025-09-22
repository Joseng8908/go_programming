package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
		}
	}
	if n.Type == html.TextNode {
		n.Data = strings.TrimSpace(n.Data)
		words += len(strings.Fields(n.Data))
	}
	// 올바른 방법
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// 재귀 호출의 결과를 새로운 변수 w, i에 받아서
		w, i := countWordsAndImages(c)
		// 기존 값에 더해줍니다.
		words += w
		images += i
	}
	return words, images
}
