package main

//go:generate statik

import (
	"fmt"
	"net/http"

	"github.com/rakyll/statik/fs"
	_ "github.com/uh-zz/traning/how_to_make_go_tool/statik_app/statik"
)

func main() {
	// Webサーバから静的コンテンツを応答するだけのアプリ
	// シングルバイナリで返せる！！

	statikFS, _ := fs.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))

	http.ListenAndServe(":8080", nil)
}
