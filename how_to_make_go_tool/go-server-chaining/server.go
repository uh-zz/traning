package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name`
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello!")
}

// レゴブロックのようにハンドラを重ねて複数のアクションを実行できる
// パイプライン処理と呼ばれる
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - ", name)
		h(w, r)
	}
}

// ユーザ権限を確認する
func protect(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p Person

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("access: ", p)
		if p.Name != "uh-zz" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/hello", protect(log(hello)))
	server.ListenAndServe()
}
