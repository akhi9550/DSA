package main

import (
	"fmt"
)

func findLUSlength(a string, b string) int {
	if a != b {
		return max(len(a), len(b))
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := "hello"
	b := "world"
	result := findLUSlength(a, b)
	fmt.Println("The length of the longest uncommon subsequence is:", result)
	a2 := "abc"
	b2 := "abc"
	result2 := findLUSlength(a2, b2)
	fmt.Println("The length of the longest uncommon subsequence is:", result2)
}
