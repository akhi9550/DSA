package main

import (
	"fmt"
)

func areaOfMaxDiagonal(dimensions [][]int) int {
	maxA, maxD := 0, 0

	for _, dim := range dimensions {
		l, w := dim[0], dim[1]
		curA := l * w
		curD := (l * l) + (w * w)

		if curD > maxD {
			maxD = curD
			maxA = curA
		} else if curD == maxD {
			if curA > maxA {
				maxA = curA
			}
		}
	}

	return maxA
}

func main() {
	dimensions := [][]int{
		{3, 4},  // Diagonal = 5, Area = 12
		{6, 8},  // Diagonal = 10, Area = 48
		{5, 12}, // Diagonal = 13, Area = 60
		{9, 12}, // Diagonal = 15, Area = 108
	}

	result := areaOfMaxDiagonal(dimensions)
	fmt.Printf("The area of the rectangle with the maximum diagonal is: %d\n", result)
}
