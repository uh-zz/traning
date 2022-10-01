package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type config struct {
	Apikey   string `json:"api_key"`
	MaxCount int    `json:"max_count"`
}

// 設定ファイルの参照プログラム
// UNIXは$HOME/.config
// Windowsは%APPDATA%配下に置くのが良いとされている
func main() {
	// os/userは環境によってコンパイル時に正しく動作しない場合があるのであまり使わない
	// u, err := user.Current()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	var configDir string

	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		configDir = os.Getenv("APPDATA")
	} else {
		configDir = filepath.Join(home, ".config")
	}

	fname := filepath.Join(configDir, ".config", "go-setting", "config.json")
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var cfg config
	err = json.NewDecoder(f).Decode(&cfg)
	fmt.Println(&cfg)
}
