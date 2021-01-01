package main

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func afterUpload(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		stream, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}

		// filepath.Base ファイル名のみ取得する
		p := filepath.Join("files", filepath.Base(header.Filename))

		f, err := os.Create(p)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}

		defer f.Close()
		io.Copy(f, stream)

		http.Redirect(w, r, path.Join("/files", p), 301)

	} else {
		http.Redirect(w, r, "/", 301)
	}
}
