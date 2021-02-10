package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Go 1.16で変更されたio/ioのReadAllの実装を理解するための試行錯誤
func main() {

	// スライスは内部に配列へのポインタを持っている
	// スライスの持っている配列を「基底配列」という
	//
	// capacityは基底配列の長さで、lengthは実際にスライスに含まれる要素数
	//
	// つまりcapacityとはスライスに割り当てるメモリ容量のこと
	//
	// capacityとスライスの要素数(length)が等しくなった場合、
	// 基底配列の長さを変更しないといけない　＝　capacityを大きくする
	b := make([]byte, 0, 512)

	fmt.Println("before len:", len(b), "cap:", cap(b))

	// ioパッケージでは以下の様に読み込む
	// Readに渡したスライスの長さだけ読み込む
	//
	// n, err := r.Read(b[len(b):cap(b)])
	//
	// nには読み込んだバイト数が入ってくる
	// b[:len(b)]には実データが入っている
	// b = b[len(b):cap(b)] -> len: 512 cap: 512
	b = b[len(b):cap(b)]

	fmt.Println("b = b[len(b):cap(b)] -> len:", len(b), "cap:", cap(b))

	// b = append(b, 0) -> len: 513 cap: 1024
	// b = append(b, 0)[:len(b)] -> len: 512 cap: 1024
	// 長さを変えずに容量だけ増やす
	// -> 実データを傷つけずにスライスの容量(基底配列の長さ)を増やす
	b = append(b, 0)[:len(b)]
	fmt.Println("b = append(b, 0) -> len:", len(b), "cap:", cap(b))

	r := strings.NewReader("hello world")

	a, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("a is", string(a))

	bi, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("b is", string(bi))
}
