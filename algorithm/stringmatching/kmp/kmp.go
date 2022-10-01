package kmp

import "fmt"

func patternMatchTable(w string) []int {
	var (
		t []int = []int{-1}
		k int
	)
	for j := 1; j < len(w); j++ {
		k = j - 1
		for w[0:k] != w[j-k:j] && k > 0 {
			k--
			s1 := w[0:k]
			s2 := w[j-k : j]
			fmt.Printf("j: %d, w[0:%d], w[%d:%d], s1: %s, s2: %s\n", j, k, j-k, j, s1, s2)
		}
		t = append(t, k)
		fmt.Printf("%+v\n", t)
	}
	return t
}

func kmp(word, text string, t []int) {
	var (
		i, j int
	)
	for i+j < len(text) {
		if word[j] == text[i+j] {
			j++
			if j == len(word) {
				// fmt.Printf("Found at %d\n", i)
				j = 0
				break
			}
		} else {
			i = i + j - t[j]
			if j > 0 {
				j = t[j]
			}
		}
	}
}

func nonkmp(word, text string) {
	var (
		i, j int
	)
	for i+j < len(text) {
		if word[j] == text[i+j] {
			j++
			if j == len(word) {
				// fmt.Printf("Found at %d\n", i)
				j = 0
				break
			}
		} else {
			i++
			j = 0
		}
	}
}
