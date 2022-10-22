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

func kmp(word, text string, patternTable []int) []int {
	if len(word) > len(text) {
		return nil
	}

	var (
		i, j    int
		matches []int
	)
	fmt.Printf("word: %s, text: %s, pattern: %v\n", word, text, patternTable)
	for i+j < len(text) {
		// fmt.Printf(" word[j]: %s, text[i+j]: %s\n", string(word[j]), string(text[i+j]))
		if word[j] == text[i+j] {
			j++
			if j == len(word) {
				fmt.Printf("Found at i: %d\n", i)
				matches = append(matches, i)

				i = i + j
				j = 0
			}
		} else {
			i = i + j - patternTable[j]
			if patternTable[j] > -1 {
				j = patternTable[j]
			} else {
				j = 0
			}
		}
	}
	return matches
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
