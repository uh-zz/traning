package main

import (
	"flag"
	"fmt"
)

var version = "1.0.0"

func main() {
	var showVersion bool

	// flagパッケージを使用すると-hでヘルプを自動生成してくれる
	// flag.BoolVar(変数のアドレス, パラメータ, デフォルト値, ヘルプの説明)
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")

	flag.Parse()
	if showVersion {
		fmt.Println("version:", version)
		return
	}
}
