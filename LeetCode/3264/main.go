package main

import (
	"fmt"
)

func getFinalState(nums []int, k int, multiplier int) []int {
	for i := 0; i < k; i++ {
		minIdx := 0
		minNum := nums[0]
		for idx, num := range nums {
			if num < minNum {
				minNum = num
				minIdx = idx
			}
		}
		nums[minIdx] *= multiplier
	}
	return nums
}

func main() {
	nums := []int{4, 2, 8, 1, 6}
	k := 3
	multiplier := 3
	result := getFinalState(nums, k, multiplier)
	fmt.Println(result)
}
