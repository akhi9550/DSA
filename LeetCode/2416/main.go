package main

import (
	"fmt"
)

func sumPrefixScores(words []string) []int {
	n := len(words)
	ret := make([]int, n)
	prefixCount := make(map[string]int)
	for _, word := range words {
		for i := 1; i <= len(word); i++ {
			prefix := word[:i]
			prefixCount[prefix]++
		}
	}
	for i, word := range words {
		for j := 1; j <= len(word); j++ {
			prefix := word[:j]
			ret[i] += prefixCount[prefix]
		}
	}

	return ret
}

func main() {
	words := []string{"abc", "ab", "bc", "b"}
	result := sumPrefixScores(words)
	fmt.Println("Prefix Scores:", result)
}
