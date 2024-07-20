package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	return append([]int{1}, digits...)
}

func main() {
	digits := []int{1, 2, 9}
	fmt.Println("Original digits:", digits)
	result := plusOne(digits)
	fmt.Println("Result after plusOne:", result)

	digits2 := []int{9, 9, 9}
	fmt.Println("Original digits:", digits2)
	result2 := plusOne(digits2)
	fmt.Println("Result after plusOne:", result2)
}
