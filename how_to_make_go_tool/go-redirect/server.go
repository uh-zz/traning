package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html><body>hello</body></html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "that's service is not found, so try another service")
}

func headerExample(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Location", "https://google.com")
	w.WriteHeader(302)
}

type Post struct {
	User   string
	Thread []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {

	post := &Post{
		User:   "uh-zz",
		Thread: []string{"hoge", "fuga", "nyao"},
	}
	json, _ := json.Marshal(post)

	w.Header().Set("Content-Type", "appcalition/json")
	w.Write(json)

}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writehead", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)

	server.ListenAndServe()
}
