package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var result int

	m := make(map[byte]bool)

	left := 0
	for right := 0; right < len(s); right++ {
		for m[s[right]] {
			delete(m, s[left])
			left++
		}

		m[s[right]] = true
		result = max(result, len(m))
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	str := "abcabcbb"
	fmt.Printf("Length of longest substring without repeating characters in \"%s\" is: %d\n", str, lengthOfLongestSubstring(str))

	str = "bbbbb"
	fmt.Printf("Length of longest substring without repeating characters in \"%s\" is: %d\n", str, lengthOfLongestSubstring(str))

	str = "pwwkew"
	fmt.Printf("Length of longest substring without repeating characters in \"%s\" is: %d\n", str, lengthOfLongestSubstring(str))
}
