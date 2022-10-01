package sort

import (
	"fmt"
)

func BubbleSort(a []int) []int {
	var count int
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			count++
			fmt.Printf("len(a)-i-1: %d, len(a): %d\n", len(a)-i-1, len(a))
			fmt.Printf("i: %d, j:%d, count: %d\n", i, j, count)
			fmt.Println(a[j], a[j+1])
			fmt.Println(a)
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
			fmt.Println(a)
		}
	}

	return a
}
