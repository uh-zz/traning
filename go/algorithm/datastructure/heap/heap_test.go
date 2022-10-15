package datastructure

import (
	"fmt"
	"testing"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	fmt.Printf("Less: h[i]: %d < h[j]: %d\n", h[i], h[j])
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	fmt.Printf("swap: h[i]: %d, h[j]: %d\n", h[i], h[j])
	h[i], h[j] = h[j], h[i]
	fmt.Printf("swap after: %v\n", h)
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestInit(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	Init(h)
	Push(h, 3)

	for h.Len() > 0 {
		fmt.Printf("Pop value: %d\n", Pop(h))
	}
}
