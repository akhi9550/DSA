package main

import (
	"fmt"
	"math"
)

func minimumSubarrayLength(nums []int, k int) int {
	sum := 0
	n := len(nums)
	var l int
	arr := make([]int, 64)
	ans := math.MaxInt
	min := func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}
	addBits := func(num int) int {
		sum := 0
		for i := 0; i < 64; i++ {
			if (num & (1 << i)) > 0 {
				arr[i]++
			}
			if arr[i] > 0 {
				sum += (1 << i)
			}
		}
		return sum
	}
	removeBits := func(num int) int {
		sum := 0
		for i := 0; i < 64; i++ {
			if (num & (1 << i)) > 0 {
				arr[i]--
			}
			if arr[i] > 0 {
				sum += (1 << i)
			}
		}
		return sum
	}

	for i := 0; i < n; i++ {
		sum = addBits(nums[i])
		for sum >= k && l <= i {
			ans = min(ans, i-l+1)
			sum = removeBits(nums[l])
			l++
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 6
	result := minimumSubarrayLength(nums, k)
	if result == -1 {
		fmt.Println("No subarray meets the criteria.")
	} else {
		fmt.Printf("The minimum length of subarray with sum >= %d is: %d\n", k, result)
	}
}
