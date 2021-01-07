#!/bin/bash

# リポジトリのタグを取得
GIT_VER=`git describe --tags`

# バイナリのビルド時に外部から変数を設定する
go build -ldflags "-X main.version=${GIT_VER}"