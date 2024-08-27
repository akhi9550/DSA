package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	matrix := make([][]int, numRows)

	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, i+1)
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = 1
		}
	}

	for i := 2; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			matrix[i][j] = matrix[i-1][j-1] + matrix[i-1][j]
		}
	}

	return matrix
}

func main() {
	numRows := 5 
	pascalTriangle := generate(numRows)

	for _, row := range pascalTriangle {
		fmt.Println(row)
	}
}
