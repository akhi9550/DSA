package main

import (
	"fmt"
)

func smallerNumbersThanCurrent(nums []int) []int {
	k := make([]int, len(nums))
	c := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				c++
			}
		}
		k[i] = c
		c = 0
	}
	return k
}

func main() {
	nums := []int{8, 1, 2, 2, 3}
	result := smallerNumbersThanCurrent(nums)
	fmt.Println("Input array:", nums)
	fmt.Println("Output array:", result)
}
