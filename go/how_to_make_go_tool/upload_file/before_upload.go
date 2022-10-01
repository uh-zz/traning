package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func beforeUpload(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		_, header, err := r.FormFile("file")
		if err != nil {

			http.NotFound(w, r)
			return
		}

		s, _ := header.Open()

		p := filepath.Join("files", header.Filename)

		buf, _ := ioutil.ReadAll(s)
		ioutil.WriteFile(p, buf, 0644)
		http.Redirect(w, r, "/upload", 301)

	} else {
		http.Redirect(w, r, "/", 301)
	}
}
