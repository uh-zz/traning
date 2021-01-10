package main

import (
	"bytes"
	"os"
)

func Readfrom() {
	f, _ := os.Open(os.Args[1])
	buf := bytes.Buffer{}

	buf.ReadFrom(f)
	buf.WriteTo(os.Stdout)
}
