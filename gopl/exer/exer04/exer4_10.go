package main

import(
	"gopl/pkg/github"
	"fmt"
	"log"
	"os"
	"time"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func issues() {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s \t %s\n", 
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `josn:"total_count`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLURL string `json:"html_url`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"created_at"`
	Body string
}

type User struct {
	Login string
	HTMLURL string `json:"html_url`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("serach query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
