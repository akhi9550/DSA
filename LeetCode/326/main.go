package main

import (
	"fmt"
)

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	if n%3 != 0 && n != 1 {
		return false
	}
	if n == 3 || n == 1 {
		return true
	}
	return isPowerOfThree(n / 3)
}

func main() {
	testValues := []int{1, 3, 9, 27, 81, 0, -3, 10, 243}

	for _, val := range testValues {
		if isPowerOfThree(val) {
			fmt.Printf("%d is a power of three.\n", val)
		} else {
			fmt.Printf("%d is not a power of three.\n", val)
		}
	}
}
