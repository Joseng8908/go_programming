package exer7

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

	// visit 함수를 호출하여 모든 링크를 추출
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit은 n에서 발견된 모든 링크를 links 슬라이스에 추가하고 그 결과를 반환합니다.
func visit(links []string, n *html.Node) []string {
	// 1. 현재 노드가 <a> 태그(ElementNode)인지 확인
	if n.Type == html.ElementNode && n.Data == "a" {
		// 2. <a> 태그의 모든 속성(Attr)을 순회
		for _, a := range n.Attr {
			// 3. 속성 중 key가 "href"인 것을 찾음
			if a.Key == "href" {
				// 4. 링크 값을 links 슬라이스에 추가
				links = append(links, a.Val)
			}
		}
	}

	// 5. 현재 노드의 자식 노드들을 순회하며 재귀적으로 visit 호출
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c) // 자식 노드로 내려가 탐색 결과를 누적
	}
	return links // 누적된 링크 목록 반환
}
