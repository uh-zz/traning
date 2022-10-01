package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Gist struct {
	Rawurl string `json:"html_url"`
}

// Doer はGistのAPIにリクエストするインターフェース
type Doer interface {
	doGistsRequest(usr string) (io.Reader, error)
}

// Client はGistのList APIを扱うためのクライアント実装
type Client struct {
	Gister Doer
}

// Gister Gist構造体
type Gister struct{}

func (g *Gister) doGistsRequest(user string) (io.Reader, error) {

	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/gists", user))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {

		return nil, err
	}
	return &buf, nil
}

// ListGists TBD
func (c *Client) ListGists(user string) ([]string, error) {

	r, err := c.Gister.doGistsRequest(user)
	if err != nil {
		return nil, err
	}

	var gists []Gist
	if err := json.NewDecoder(r).Decode(&gists); err != nil {
		return nil, err
	}

	urls := make([]string, 0, len(gists))
	for _, u := range gists {
		urls = append(urls, u.Rawurl)
	}
	return urls, nil
}

// func main() {
// 	urls, err := ListGists("mattn")
// 	if err != nil {
// 		fmt.Println("error", err)
// 	}

// 	for i, u := range urls {
// 		if i > 10 {
// 			return
// 		}
// 		fmt.Println(u)
// 	}
// }
