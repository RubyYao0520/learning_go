package main

//2026/1/1
//各种数据结构

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssueURL = "https://api.github.com/search/issues"

type IssuesSearchResults struct {
	TotalCount int
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"tota;_count"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func Day007() {
	fmt.Print("Hello")

}

// 参考GO语言圣经，获取GitHub Issue的一个函数
func Search(terms []string) (*IssuesSearchResults, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	//获取响应和错误
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query faild: %d", resp.Status)
	}

	var result IssuesSearchResults
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
