package main

import (
	"fmt"
)

func minLength(s string) int {
	i := 0
	for i < len(s) {
		n := s[max(0, i-1):i+1]
		if n == "AB" || n == "CD" {
			s = s[:i-1] + s[i+1:]
			i = i - 1
		} else {
			i++
		}
	}
	return len(s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var input string
	fmt.Println("Enter a string:")
	fmt.Scanln(&input)
	result := minLength(input)
	fmt.Printf("The length of the string after removing 'AB' or 'CD' is: %d\n", result)
}
