package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// httpリクエストは論理パスなのでpathパッケージを使う
		if ok, err := path.Match("/data/*.html", r.URL.Path); err != nil || !ok {

			http.NotFound(w, r)
			return
		}
		// 以降はパスを物理パスとして使うのでpath/filepathパッケージを使う
		name := filepath.Join(cwd, "data", filepath.Base(r.URL.Path))

		f, err := os.Open(name)
		if err != nil {
			http.NotFound(w, r)
		}

		defer f.Close()
		io.Copy(w, f)
	})

	http.ListenAndServe(":8080", nil)

}
