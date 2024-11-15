package main

import (
	"fmt"
)

func reversePrefix(word string, ch byte) string {
	a := []byte(word)
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			for j := 0; j <= i/2; j++ {
				temp := a[j]
				a[j] = a[i-j]
				a[i-j] = temp
			}
			return string(a)
		}
	}
	return word
}

func main() {
	word := "abcdefd"
	ch := 'd'
	fmt.Println(reversePrefix(word, byte(ch)))
}
