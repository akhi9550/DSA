package main

import (
	"fmt"
	"math"
)

func findComplement(num int) int {
	var res []int
	for num > 0 {
		rem := num % 2
		num /= 2
		res = append(res, rem^1)
	}

	var complement float64
	for i, v := range res {
		complement += math.Pow(2, float64(i)) * float64(v)
	}

	return int(complement)
}

func main() {
	tests := []int{5, 1, 7, 10}
	for _, test := range tests {
		fmt.Printf("Number: %d, Complement: %d\n", test, findComplement(test))
	}
}
