package main

import (
	"fmt"
	"math"
)

func maxMatrixSum(matrix [][]int) int64 {
	sum := int64(0)
	minValue, nNeg := math.MaxInt32, 0

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] < 0 {
				nNeg++
				sum -= int64(matrix[i][j])
				minValue = min(minValue, -matrix[i][j])
			} else {
				sum += int64(matrix[i][j])
				minValue = min(minValue, matrix[i][j])
			}
		}
	}

	if nNeg%2 == 1 {
		sum -= int64(2 * minValue)
	}

	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	matrix := [][]int{
		{1, -2, 3},
		{-4, 5, -6},
		{7, -8, 9},
	}
	result := maxMatrixSum(matrix)
	fmt.Println("Maximum Matrix Sum:", result)
}
