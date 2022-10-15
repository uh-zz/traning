package datastructure

import (
	"fmt"
	"sort"
)

// FYI: https://cs.opensource.google/go/go/+/refs/tags/go1.19.2:src/container/heap/heap.go
type Interface interface {
	sort.Interface
	Push(x any)
	Pop() any
}

func Init(h Interface) {
	fmt.Println("Init start")
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		fmt.Printf("Init: i: %v\n", i)
		down(h, i, n)
	}
	fmt.Println("Init end")
}

func Push(h Interface, x any) {
	fmt.Println("Push start")
	h.Push(x)
	up(h, h.Len()-1)
	fmt.Println("Push end")
}

func Pop(h Interface) any {
	defer fmt.Println("Pop end")
	fmt.Println("Pop start")
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func up(h Interface, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down(h Interface, i0, n int) {
	fmt.Println("down start")
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1
		fmt.Printf("down: j: %v\n", j)
		if j2 := j1 + 1; j2 < n && !h.Less(j1, j2) {
			j = j2
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	fmt.Println("down end")
}
