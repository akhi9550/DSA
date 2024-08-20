package main

import "fmt"

func countBits(n int) []int {
	results := make([]int, n+1)
	for i := 0; i <= n; i++ {
		results[i] = countBitsOf(i)
	}
	return results
}

func countBitsOf(v int) int {
	count := 0
	for v > 0 {
		if v%2 != 0 {
			count++
		}
		v /= 2
	}
	return count
}

func main() {
	n := 5
	result := countBits(n)
	fmt.Printf("Number of 1s in binary representation from 0 to %d: %v\n", n, result)
}
