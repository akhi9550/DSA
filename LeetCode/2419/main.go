package main

import (
	"fmt"
)

func longestSubarray(nums []int) int {
	result, max, length := 0, 0, 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	for _, v := range nums {
		if v == max {
			length++
			if length > result {
				result = length
			}
		} else {
			length = 0
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 3, 2, 2, 4, 4, 4, 4}
	result := longestSubarray(nums)
	fmt.Println("The length of the longest subarray with the maximum element is:", result)
}
