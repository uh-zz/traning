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
	// vもしくはversionという文字列をパラメータにしてflag.Parseすると、
	// v/version -> trueと解釈される
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")

	// アドレスを渡しただけなのでまだfalse
	fmt.Println(showVersion)

	flag.Parse()

	// Parse後のフラグはtrue
	if showVersion {

		checkVersion()

		fmt.Println("version:", version)
		return
	}
}
