#!/bin/bash

# リポジトリのタグを取得
GIT_VER=`git describe --tags`

# バイナリのビルド時に外部から変数を設定する
go build -ldflags "-X main.version=${GIT_VER}"

# main以外のパッケージの変数を設定する場合はimport pathで指定
# go build -ldflags "-X github.com/uh-zz/hogehoge.Version=${GIT_VER}"