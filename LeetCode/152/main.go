package main

import (
	"fmt"
)

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := make([]int, len(nums))
	min := make([]int, len(nums))
	max[0] = nums[0]
	min[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			max[i], min[i] = 0, 0
			continue
		}
		if max[i-1] > 0 && min[i-1] > 0 {
			if nums[i] > 0 {
				max[i] = max[i-1] * nums[i]
				min[i] = nums[i]
			} else {
				max[i] = nums[i]
				min[i] = max[i-1] * nums[i]
			}
		}
		if max[i-1] > 0 && min[i-1] < 0 {
			if nums[i] > 0 {
				max[i] = max[i-1] * nums[i]
				min[i] = min[i-1] * nums[i]
			} else {
				max[i] = min[i-1] * nums[i]
				min[i] = max[i-1] * nums[i]
			}
		}
		if max[i-1] < 0 && min[i-1] < 0 {
			if nums[i] > 0 {
				max[i] = nums[i]
				min[i] = min[i-1] * nums[i]
			} else {
				max[i] = min[i-1] * nums[i]
				min[i] = nums[i]
			}
		}
		if max[i-1] == 0 && min[i-1] == 0 {
			max[i], min[i] = nums[i], nums[i]
		}
	}
	result := max[0]
	for i := 1; i < len(max); i++ {
		if max[i] > result {
			result = max[i]
		}
	}
	return result
}

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println("Maximum product subarray:", maxProduct(nums))

	nums = []int{-2, 0, -1}
	fmt.Println("Maximum product subarray:", maxProduct(nums))

	nums = []int{0, 2}
	fmt.Println("Maximum product subarray:", maxProduct(nums))

	nums = []int{2, -5, -2, -4, 3}
	fmt.Println("Maximum product subarray:", maxProduct(nums))
}
