package datastructure

type Node struct {
	data int
	next *Node
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) SetData(data int) {
	n.data = data
}

func (n *Node) SetNext(node *Node) {
	n.next = node
}
