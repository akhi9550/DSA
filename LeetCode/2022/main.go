package main

import (
	"fmt"
)

func construct2DArray(original []int, m int, n int) [][]int {
	matrix := [][]int{}
	if len(original) != n*m {
		return matrix
	}

	for x := 0; x < m; x++ {
		matrix = append(matrix, original[0:n])
		original = original[n:]
	}
	return matrix
}

func main() {
	original := []int{1, 2, 3, 4, 5, 6}
	m, n := 2, 3

	result := construct2DArray(original, m, n)

	fmt.Println("2D Array:")
	for _, row := range result {
		fmt.Println(row)
	}
}
