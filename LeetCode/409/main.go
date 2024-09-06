package main

import (
	"fmt"
)

func longestPalindrome(s string) int {
	count := make([]int, 128)
	result := 0
	
	for _, item := range s {
		count[item]++
	}
	for _, item := range count {
		result += (item / 2) * 2
		if result%2 == 0 && item%2 == 1 {
			result++
		}
	}

	return result
}

func main() {
	testStrings := []string{"abccccdd", "a", "bb", "abc", "abbcc"}

	for _, s := range testStrings {
		fmt.Printf("The longest palindrome length that can be built from \"%s\" is: %d\n", s, longestPalindrome(s))
	}
}
