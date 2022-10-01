package main

// cgoでは各環境により異なるCFLAGSやLIBSを個別に記述する
// #cgo CFLAGS: -DPNG_DEBUG=1
// #cgo amd64 386 CFLAGS: -DX86=1
// #cgo LDFLAGS: -lpng
// #include <png.h>

// 上記のように指定するとライブラリがバージョンにより異なるライブラリファイル名を提供している
// 場合に対応できない

// cgoではpkg-configを扱えるようになっている
// pkg-configでパッケージのコンパイルオプションを取得できる

// 以下のように記述することでgo buildを実行時に必要なライブラリをリンクできるようになる
// #cgo pkg-config: png cairo
// #include <png.h>

import "C"
