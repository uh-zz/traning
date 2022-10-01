package main

import (
	"fmt"
	"sync"
)

func main() {

	var mux sync.Mutex

	c := make(chan bool)

	m := make(map[string]string)

	// 無名関数はいつ実行されるかわからない
	go func() {
		// スレッドセーフにするために必要
		mux.Lock()

		m["1"] = "a"

		mux.Unlock()
		c <- true

	}()

	// スレッドセーフにするために必要
	mux.Lock()

	// 大事なのは、mの同じアドレスにm["1"] = "a"を書き込もうとしていること
	m["2"] = "b"

	mux.Unlock()

	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
