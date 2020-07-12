package main

import (
	"fmt"
	"time"

	"github.com/uh-zz/traning/algorithm/shuffle"
)

func main() {

	input := shuffle.RandomIntList(1000000)

	startTime := time.Now()

	inputLength := len(input)

	// fmt.Println("before sort->", input)

	MergeSort(&input, 0, inputLength)

	// fmt.Println("after sort->", input)

	endTime := time.Now()
	fmt.Printf("%f seconds\n", (endTime.Sub(startTime)).Seconds())
}

// MergeSort マージソート
func MergeSort(input *[]int, left, right int) {

	// 要素数1
	if right-left == 1 {
		return
	}

	middle := left + (right-left)/2

	// 配列の左側
	// fmt.Println("left split", "index[from]", left, "index[to]", middle)
	MergeSort(input, left, middle)
	// 配列の右側
	// fmt.Println("right split", "index[from]", middle, "index[to]", right)
	MergeSort(input, middle, right)

	var buffer []int

	// 左側と右側をバッファにためる（右側反転）
	for index := left; index < middle; index++ {
		buffer = append(buffer, (*input)[index])
	}
	for index := right - 1; index >= middle; index-- {
		buffer = append(buffer, (*input)[index])
	}

	// fmt.Println("buffer->", buffer)
	// fmt.Println("[calculating] before input->", input)
	// マージする
	scopeLeft := 0
	scopeRight := len(buffer) - 1
	for index := left; index < right; index++ {

		if buffer[scopeLeft] <= buffer[scopeRight] {
			// 左側採用
			(*input)[index] = buffer[scopeLeft]
			scopeLeft++
		} else {
			// 右側採用
			(*input)[index] = buffer[scopeRight]
			scopeRight--
		}
	}
	// fmt.Println("[calculating] after input->", input)
}
