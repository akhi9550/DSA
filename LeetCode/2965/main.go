package main

import (
	"fmt"
)

func findMissingAndRepeatedValues(grid [][]int) []int {
	a := 0
	b := 0
	n := len(grid)
	table := make([]int, n*n+1)

	for _, row := range grid {
		for _, val := range row {
			table[val]++
		}
	}

	for i := 1; i <= n*n; i++ {
		if table[i] == 2 {
			a = i
		}
		if table[i] == 0 {
			b = i
		}
	}

	return []int{a, b}
}

func main() {
	grid := [][]int{
		{1, 2, 3},
		{4, 6, 6},
		{7, 8, 9},
	}

	result := findMissingAndRepeatedValues(grid)
	fmt.Printf("Repeated Value: %d\n", result[0])
	fmt.Printf("Missing Value: %d\n", result[1])
}
