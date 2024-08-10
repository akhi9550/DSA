package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	n := len(nums)
	res := 0
	for i, j := 0, 1; j <= n; j++ {
		res = res + j
		res = res - nums[i]
		i++
	}
	return res
}

func main() {
	nums := []int{3, 0, 1}
	fmt.Println("Missing number is:", missingNumber(nums)) // Output: 2

	nums = []int{9,6,4,2,3,5,7,0,1}
	fmt.Println("Missing number is:", missingNumber(nums)) // Output: 8

	nums = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	fmt.Println("Missing number is:", missingNumber(nums)) // Output: 10
}
