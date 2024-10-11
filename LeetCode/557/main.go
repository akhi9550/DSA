package main

import (
	"fmt"
)

func reverseWords(s string) string {
	res := []rune(s)

	swap := func(i, j int) {
		for ; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}

	f, l := 0, 1
	for l < len(res) {
		if res[l] == ' ' {
			swap(f, l-1)
			f = l + 1
			l = f + 1
		} else {
			l++
		}
	}

	swap(f, l-1)

	return string(res)
}

func main() {
	s := "hello world this is Golang"
	result := reverseWords(s)
	fmt.Println("Original String: ", s)
	fmt.Println("Reversed Words:  ", result)
}
