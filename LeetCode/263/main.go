package main

import (
	"fmt"
)

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	return n == 1
}

func main() {
	testNumbers := []int{6, 8, 14, 1, 30, -5}

	for _, num := range testNumbers {
		if isUgly(num) {
			fmt.Printf("%d is an ugly number.\n", num)
		} else {
			fmt.Printf("%d is not an ugly number.\n", num)
		}
	}
}
