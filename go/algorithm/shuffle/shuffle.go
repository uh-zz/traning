package shuffle

import (
	"math/rand"
)

// RandomIntList ランダムなintリスト取得
func RandomIntList(n int) []int {

	intSlice := make([]int, n)
	for index := 0; index < n; index++ {
		intSlice[index] = index
	}

	shuffle(&intSlice)

	return intSlice
}

func shuffle(array *[]int) {

	for index := len(*array) - 1; index >= 0; index-- {
		random := rand.Intn(index + 1)
		(*array)[index], (*array)[random] = (*array)[random], (*array)[index]
	}
}
