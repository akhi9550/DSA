package main

import (
	"fmt"
)

func largeGroupPositions(s string) [][]int {
	ans := make([][]int, 0)
	ind := 0
	for i := 0; i <= len(s); i++ {
		if i == 0 || i == len(s) || s[ind] != s[i] {
			if i-ind >= 3 {
				ans = append(ans, []int{ind, i - 1})
			}
			ind = i
		}
	}
	return ans
}

func main() {
	tests := []string{
		"abbxxxxzzy",
		"abc",
		"aaa",
		"aaabbbcccddd",
		"",
	}

	for _, test := range tests {
		fmt.Printf("Input: %q\n", test)
		fmt.Printf("Output: %v\n\n", largeGroupPositions(test))
	}
}
