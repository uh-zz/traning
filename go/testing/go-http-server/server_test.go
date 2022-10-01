package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {

	// テスト用のサーバを立ち上げる
	ts := httptest.NewServer(Route())
	defer ts.Close()

	// 通常通りリクエストを遅れる
	// テストサーバURLはts.URLで取得
	res, err := http.Get(ts.URL + "/greet?name=gopher")
	if err != nil {
		t.Fatalf("http.Get failed: %s", err)
	}

	defer res.Body.Close()
	greeting, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read from http response body failed: %s", err)
	}

	expected := "hello, phper"
	if string(greeting) != expected {
		t.Fatalf("response of /greet?name=gopher returns %s, want %s",
			string(greeting), expected)
	}
}
