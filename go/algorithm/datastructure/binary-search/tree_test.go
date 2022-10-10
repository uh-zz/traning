package datastructure

import (
	"testing"
)

func TestHoge(t *testing.T) {
	// データ作成
	data := []int{5, 3, 7, 1, 4, 6, 8}

	// 木の作成
	tree := NewTree()
	tree.root = NewNode(0, 0)

	// 二分探索木の作成
	for _, v := range data {
		tree.Insert(tree.root, v, v)
	}

	// 木の中身を表示
	tree.Traverse(tree.root)

	// 二分探索木の検索
	_ = tree.Search(tree.root, 1)
	_ = tree.Search(tree.root, 10)
}
