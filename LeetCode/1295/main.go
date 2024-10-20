package main

import (
	"fmt"
	"strconv"
)

func findNumbers(nums []int) int {
	res := 0
	for _, val := range nums {
		temp := strconv.Itoa(val)
		if len(temp)%2 == 0 {
			res++
		}
	}
	return res
}

func main() {
	nums := []int{12, 345, 2, 6, 7896}
	result := findNumbers(nums)
	fmt.Printf("Numbers with an even number of digits: %d\n", result)
}
