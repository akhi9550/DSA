package main

import (
	"fmt"
)

func resultsArray(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	n := len(nums)
	l := 0
	cnt := 1
	for r := 1; r < n; r++ {
		if nums[r] == nums[r-1]+1 {
			cnt++
		} else {
			cnt = 1
		}

		if cnt == k {
			for ; l < r-k+1; l++ {
				nums[l] = -1
			}
			nums[l] = nums[r]
			l++
			cnt--
		}
	}
	for ; l < n-k+1; l++ {
		nums[l] = -1
	}

	return nums[:n-k+1]
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := 3
	result := resultsArray(nums, k)
	fmt.Println("Result:", result)
}
