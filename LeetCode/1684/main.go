package main

import (
	"fmt"
)

func countConsistentStrings(allowed string, words []string) int {
	ans, tmp := 0, make([]int, 26)
	for _, v := range allowed {
		tmp[v-'a'] = 1
	}
	for _, v := range words {
		b := true
		for _, c := range v {
			if tmp[c-'a'] == 0 {
				b = false
				break
			}
		}
		if b {
			ans++
		}
	}
	return ans
}

func main() {
	allowed := "ab"
	words := []string{"ad", "bd", "aaab", "baa", "badab"}

	result := countConsistentStrings(allowed, words)

	fmt.Printf("Number of consistent strings: %d\n", result)
}
