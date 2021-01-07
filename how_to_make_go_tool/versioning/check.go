package main

import (
	"fmt"

	"github.com/tcnksm/go-latest"
)

// go-latestを組み込むことでバージョンの管理ができる
func checkVersion() {
	githubTag := &latest.GithubTag{
		Owner:      "tcnksm",
		Repository: "ghr",
	}

	res, _ := latest.Check(githubTag, "0.1.0")
	if res.Outdated {
		fmt.Printf("0.1.0 is not latest, you should update to %s", res.Current)
	}
}
