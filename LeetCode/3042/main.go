package main

import (
	"fmt"
	"strings"
)

func countPrefixSuffixPairs(words []string) int {
	c := 0
	for i, v := range words {
		for _, nestedWord := range words[i+1:] {
			isPre := strings.HasPrefix(nestedWord, v)
			isSuf := strings.HasSuffix(nestedWord, v)
			if isPre && isSuf {
				c++
			}
		}
	}
	return c
}

func main() {
	words := []string{"abc", "ab", "bc", "abcab", "bcbc"}
	fmt.Println("Count of prefix-suffix pairs:", countPrefixSuffixPairs(words))
}
