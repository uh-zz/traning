package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func getHTTP(url string, dst io.Writer) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)

	// contextを渡したリクエスト
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		// レスポンスヘッダ取得までに10秒経過した場合はエラー
		return err
	}

	defer resp.Body.Close()

	_, err := io.Copy(dst, resp.Body)

	// ボディ取得完了までに10秒経過した場合はエラー
	return err

}
