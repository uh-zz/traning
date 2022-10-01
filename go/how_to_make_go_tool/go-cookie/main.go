package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("hello, world")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg), // ヘッダ内のクッキーはURLエンコードする必要がらある
	}
	http.SetCookie(w, &c)
}

// フラッシュメッセージの実装
func showMessage(w http.ResponseWriter, r *http.Request) {

	// すでに格納されているクッキーを取得
	c, err := r.Cookie("flash")

	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "cookie has no message")
			return
		}
	}

	// 同じ名前のクッキーで上書き
	// 期間指定を過去にしているので、新しく作ったこのクッキーも削除される
	rc := http.Cookie{
		Name:    "flash",
		MaxAge:  -1,              // 負の値
		Expires: time.Unix(1, 0), // 過去日
	}
	http.SetCookie(w, &rc)

	val, _ := base64.URLEncoding.DecodeString(c.Value)
	fmt.Fprintln(w, string(val))
}

func setCookie(w http.ResponseWriter, r *http.Request) {

	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "hoge",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "fuga",
		HttpOnly: true,
	}

	// w.Header().Set("Set-Cookie", c1.String())
	// // w.Header().Add()でクッキーを追加
	// w.Header().Add("Set-Cookie", c2.String())

	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Println(w, h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)

	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)

	server.ListenAndServe()
}
