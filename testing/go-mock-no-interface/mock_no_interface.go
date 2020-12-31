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

var doGistsRequest = func(user string) (io.Reader, error) {

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

func ListGists(user string) ([]string, error) {

	r, err := doGistsRequest(user)
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
