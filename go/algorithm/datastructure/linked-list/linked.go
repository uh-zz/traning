package datastructure

import "fmt"

type LinkedList struct {
	root *Node
	len  int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Insert(current *Node, data int) *Node {
	node := NewNode()
	node.SetData(data)

	// when Linked list is empty, current is not associated with any node
	if l.len == 0 {
		l.root = node
		current = l.root
	} else {
		fmt.Printf("current: %v\n", current)
		node.SetNext(current.next)
		current.SetNext(node)
		current = current.next
	}
	l.len++
	return current
}

func (l *LinkedList) Traverse() {
	node := l.root
	for node != nil {
		println(node.data)
		node = node.next
	}
}
