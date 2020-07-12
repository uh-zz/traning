package main

import (
	"fmt"
	"time"

	"github.com/uh-zz/traning/algorithm/shuffle"
)

func main() {

	input := shuffle.RandomIntList(1000000)

	startTime := time.Now()

	// 挿入ソート
	// fmt.Println("start input->", input)
	for index := 1; index < len(input); index++ {

		// 挿入したい値
		insertValue := input[index]
		// fmt.Println("insertValue->", insertValue)

		// 挿入位置
		point := index
		for ; point > 0; point-- {
			if input[point-1] > insertValue {
				input[point] = input[point-1]
			} else {
				break
			}
		}
		input[point] = insertValue
		// fmt.Println("[calculating] input->", input)
	}
	// fmt.Println("end input->", input)

	endTime := time.Now()
	fmt.Printf("%f seconds\n", (endTime.Sub(startTime)).Seconds())
}
