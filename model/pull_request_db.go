package model

import (
	"fmt"
)

type PullRequestDB struct {
	Id         int    `json:"id"`
	URL        string `json:"url"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	User       string `json:"user"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
	Merged     string `json:"merged"`
	Closed     string `json:"closed"`
	Screenshot string `json:"screenshot"`
	RawData    string `json:"rawdata"`
}

func PullRequestDBFromGitHubEvent(ghe map[string]interface{}) PullRequestDB {
	var ret PullRequestDB

	if id, ok := ghe["id"]; ok {
		ret.Id = int(id.(float64))
	}
	if html_url, ok := ghe["html_url"]; ok {
		ret.URL = fmt.Sprintf("%v", html_url)
	}
	if title, ok := ghe["title"]; ok {
		ret.Title = fmt.Sprintf("%v", title)
	}
	if body, ok := ghe["body"]; ok {
		ret.Body = fmt.Sprintf("%v", body)
	}
	if val, ok := ghe["user"]; ok {
		user := val.(map[string]interface{})
		if login, ok := user["login"]; ok {
			ret.User = fmt.Sprintf("%v", login)
		}
	}
	if created, ok := ghe["created_at"]; ok {
		ret.Created = fmt.Sprintf("%v", created)
	}
	if updated, ok := ghe["updated_at"]; ok {
		ret.Updated = fmt.Sprintf("%v", updated)
	}
	if merged, ok := ghe["merged_at"]; ok {
		ret.Merged = fmt.Sprintf("%v", merged)
	}
	if closed, ok := ghe["closed_at"]; ok {
		ret.Closed = fmt.Sprintf("%v", closed)
	}
	if url, ok := ghe["url"]; ok {
		ret.RawData = fmt.Sprintf("%v", url)
	}

	return ret
}
