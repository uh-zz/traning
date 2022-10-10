package datastructure

type Node struct {
	key   int
	value int
	left  *Node
	right *Node
}

func NewNode(key, value int) *Node {
	return &Node{key: key, value: value}
}
