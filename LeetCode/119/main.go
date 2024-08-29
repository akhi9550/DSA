package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	res := append(getRow(rowIndex-1), 1)
	for i := rowIndex - 1; i > 0; i-- {
		res[i] = res[i-1] + res[i]
	}
	return res
}

func main() {
	rowIndex := 5 
	row := getRow(rowIndex)
	fmt.Printf("Row %d of Pascal's Triangle: %v\n", rowIndex, row)
}
