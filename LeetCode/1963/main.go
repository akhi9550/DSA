package main

import (
	"fmt"
)

func minSwaps(s string) int {
	open, close := 0, 0
	swaps := 0

	for _, ch := range s {
		if ch == '[' {
			open++
		} else {
			close++
		}

		if close > open {
			swaps++
			close--
			open++
		}
	}
	return swaps
}

func main() {
	s := "[]][]["
	fmt.Println("Input:", s)
	result := minSwaps(s)
	fmt.Println("Minimum number of swaps required:", result)

	s = "[][[["
	fmt.Println("\nInput:", s)
	result = minSwaps(s)
	fmt.Println("Minimum number of swaps required:", result)
}
