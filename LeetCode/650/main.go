package main

import (
	"fmt"
)

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			return minSteps(n/i) + i
		}
	}
	return -1
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	steps := minSteps(n)
	fmt.Printf("Minimum steps to get %d is: %d\n", n, steps)
}
