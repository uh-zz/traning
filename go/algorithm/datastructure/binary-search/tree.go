package datastructure

import "fmt"

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(parent *Node, key, value int) {
	fmt.Printf("insert parent->%+v, key->%d, value->%d\n", parent, key, value)

	if parent == nil {
		t.root = NewNode(key, value)
		return
	}
	child := NewNode(key, value)
	if child.value >= parent.value {
		if parent.right == nil {
			parent.right = child
		} else {
			t.Insert(parent.right, key, value)
		}
	} else {
		if parent.left == nil {
			parent.left = child
		} else {
			t.Insert(parent.left, key, value)
		}
	}
	fmt.Printf("end insert parent->%+v\n", parent)
}

func (t *Tree) Search(parent *Node, key int) *Node {
	if parent == nil {
		fmt.Printf("not found key->%d\n", key)
		return nil
	}
	if parent.key == key {
		fmt.Printf("found parent->%+v\n", parent)
		return parent
	}
	if key >= parent.key {
		fmt.Printf("search right parent->%+v\n", parent)
		return t.Search(parent.right, key)
	}
	fmt.Printf("search left parent->%+v\n", parent)
	return t.Search(parent.left, key)
}

func (t *Tree) Traverse(parent *Node) {
	if parent == nil {
		return
	}
	t.Traverse(parent.left)
	fmt.Printf("parent->%+v\n", parent)
	t.Traverse(parent.right)
}
