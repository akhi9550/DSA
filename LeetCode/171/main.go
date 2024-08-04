package main

import (
	"fmt"
	"math"
)

func titleToNumber(columnTitle string) int {
	res := 0

	for i := len(columnTitle) - 1; i >= 0; i-- {
		res += (int(columnTitle[i]) - 64) * int(math.Pow(26, float64(len(columnTitle)-1-i)))
	}

	return res
}

func main() {
	tests := []string{"A", "B", "C", "Z", "AA", "AB", "ZY"}

	for _, test := range tests {
		fmt.Printf("Column Title: %s -> Column Number: %d\n", test, titleToNumber(test))
	}
}
