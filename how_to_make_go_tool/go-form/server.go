package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// fmt.Fprintln(w, r.PostForm)

	file, _, err := r.FormFile("uploaded")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, string(data))
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
