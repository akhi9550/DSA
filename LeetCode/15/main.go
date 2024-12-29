package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}
	var ret [][]int
	sort.Ints(nums)
	
	for i := 0; i < len(nums)-2; i++ {
		cur_target := 0 - nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left]+nums[right] < cur_target {
				left++
			} else if nums[left]+nums[right] > cur_target {
				right--
			} else {
				resub := make([]int, 3)
				resub[0] = nums[i]
				resub[1] = nums[left]
				resub[2] = nums[right]
				ret = append(ret, resub)
				for left < right && nums[left] == resub[1] {
					left++
				}
				for left < right && nums[right] == resub[2] {
					right--
				}
			}
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ret
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)
}
