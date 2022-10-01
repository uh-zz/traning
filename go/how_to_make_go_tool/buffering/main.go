package main

func main() {
	// bytes.Bufferを用いたファイル読み込み
	// 一気に読み込んで一気に出力するパターン
	Readfrom()

	// bufioを用いたファイル読み込み
	// I/Oのバッファリング
	// 省メモリなのはこっち
	Bufio()

	Checktty()
}
