package sort

import "fmt"

func OtherMergeSort(arr []int, left, right int) {
	fmt.Printf("call: arr->%+v, left->%d, right->%d\n", arr, left, right)

	if left < right {
		middle := left + (right-left)/2
		OtherMergeSort(arr, left, middle)
		OtherMergeSort(arr, middle+1, right)
		OtherMerge(arr, left, middle, right)
	}
}

func OtherMerge(arr []int, left, middle, right int) {
	leftArr := make([]int, middle-left+1)
	rightArr := make([]int, right-middle)

	for i := 0; i < len(leftArr); i++ {
		leftArr[i] = arr[left+i]
	}

	for i := 0; i < len(rightArr); i++ {
		rightArr[i] = arr[middle+1+i]
	}

	fmt.Printf("left->%d, middle->%d, right->%d,leftArr->%+v, rightArr->%+v\n", left, middle, right, leftArr, rightArr)

	var (
		leftIndex  int
		rightIndex int
	)
	for leftIndex < len(leftArr) && rightIndex < len(rightArr) {
		if leftArr[leftIndex] <= rightArr[rightIndex] {
			arr[left] = leftArr[leftIndex]
			leftIndex++
		} else {
			arr[left] = rightArr[rightIndex]
			rightIndex++
		}
		left++
	}

	fmt.Printf("leftArr->%+v, rightArr->%+v, arr->%+v, leftIndex->%d, rightIndex->%d, left->%d\n", leftArr, rightArr, arr, leftIndex, rightIndex, left)

	for leftIndex < len(leftArr) {
		arr[left] = leftArr[leftIndex]
		leftIndex++
		left++
	}

	for rightIndex < len(rightArr) {
		arr[left] = rightArr[rightIndex]
		rightIndex++
		left++
	}

	fmt.Printf("arr->%+v\n", arr)
}
