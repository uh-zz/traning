package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	// Go 1.16の検証
	//
	// シグナル付きのコンテキスト
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	go func() {
		// close要求を受け取るまでブロック
		<-ctx.Done()

		// タイムアウトつきコンテキスト作成
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// サーバのgraceful shutdown
		//
		// shutdown中でも5秒たったら、強制的にタイムアウトさせる
		server.Shutdown(ctx)
	}()

	// サーバ起動
	log.Fatal(server.ListenAndServe())
}
