package main

import "fmt"

func main() {
	result := fibo(6)
	fmt.Println(result)

	f := fiboClosure()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d,", f())
	}
}

func fibo(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	// fmt.Println("fibo(", n-2, ") + fibo(", n-1, ")")
	return fibo(n-2) + fibo(n-1)
}

func fiboClosure() func() int {
	x, y := 0, 1
	return func() (result int) {
		result = x
		x, y = y, x+y
		return
	}
}
