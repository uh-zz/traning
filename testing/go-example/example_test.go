package calc

import "fmt"

func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}

// 標準出力は
// 1
// 2
// 3
// 、、だけど、Unordered ouputのため、テストはパス！
func ExampleUnordered() {
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}
	// Unordered output:
	// 2
	// 3
	// 1
}

func ExampleShufullWillBeFailed() {
	x := map[string]int{"a": 1, "b": 2, "c": 3}

	for k, v := range x {
		fmt.Printf("k=%s v=%d\n", k, v)
	}

	// Output:
	// k=a v=1
	// k=b v=2
	// k=c v=3
}
