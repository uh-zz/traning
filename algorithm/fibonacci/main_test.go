package main

import (
	"fmt"
	"testing"
)

func TestFibo(t *testing.T) {
	expected := 8
	result := fibo(6)
	if expected != result {
		t.Fatalf("fibo(6) want %d, got %d", expected, result)
	}
}

func TestFiboClosure(t *testing.T) {
	f := fiboClosure()
	for i := 0; i < 10; i++ {
		res1 := fibo(i)
		res2 := f()
		fmt.Println("res1:", res1, "res2:", res2)
	}
}

func BenchmarkFibo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 20; i++ {
		fibo(i)
	}
}

func BenchmarkFiboClosure(b *testing.B) {

	f := fiboClosure()
	b.ResetTimer()
	for i := 0; i < 1000; i++ {
		f()
	}
}
