package exer5

import "golang.org/x/net/html"

func main() {

}
func visit5_5(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			links = addHref(links, n)
		case "link":
			links = addSrc(links, n)
		case "script":
			links = addSrc(links, n)
		case "img":
			links = addHref(links, n)
		}
	}

	links = visit5_5(links, n.FirstChild)
	links = visit5_5(links, n.NextSibling)
	return links
}

func addHref(links []string, n *html.Node) []string {
	for _, a := range n.Attr {
		if a.Key == "href" {
			links = append(links, a.Val)
		}
	}
	return links
}

func addSrc(links []string, n *html.Node) []string {
	for _, a := range n.Attr {
		if a.Key == "src" {
			links = append(links, a.Val)
		}
	}
	return links
}
