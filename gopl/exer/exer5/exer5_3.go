package exer5

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	printTextNode(doc)

}
func printTextNode(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}
	if n.Type == html.TextNode {
		data := strings.TrimSpace(n.Data)
		if data != "" {
			fmt.Println(data)
		}
	}

	printTextNode(n.FirstChild)
	printTextNode(n.NextSibling)
	return
}
