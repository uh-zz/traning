package main

import (
	"bufio"
	"os"
)

func Bufio() {
	f, _ := os.Open(os.Args[1])
	rb := bufio.NewReaderSize(f, 5)
	wb := bufio.NewWriterSize(os.Stdout, 5)

	for {
		n, _ := rb.WriteTo(wb)
		if n == 0 {
			break
		}
	}
	wb.Flush()
}
