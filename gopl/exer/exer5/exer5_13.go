package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	parsedURL, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	hostname := parsedURL.Hostname()
	crawl := func(item string) []string {
		fmt.Println(item)
		list, err := Extract(item)
		var filteredList []string
		if err != nil {
			log.Print(err)
		}
		for _, link := range list {
			linkURL, err := url.Parse(link)
			if err != nil {
				log.Print(err)
			}
			if linkURL.Hostname() == hostname {
				filteredList = append(list, link)
			}
		}
		return filteredList
	}

	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)

	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)

			}
		}
	}
}
