package fibonacci

func fibo(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	// fmt.Println("fibo(", n-2, ") + fibo(", n-1, ")")
	return fibo(n-2) + fibo(n-1)
}
