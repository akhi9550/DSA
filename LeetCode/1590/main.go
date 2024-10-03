package main

import "fmt"

func minSubarray(nums []int, p int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	remainder := totalSum % p
	if remainder == 0 {
		return 0 
	}
	subarraySum := 0
	minLength := len(nums)
	lastOccurrence := make(map[int]int)
	lastOccurrence[0] = -1

	for i, num := range nums {
		subarraySum = (subarraySum + num) % p
		neededRemainder := (subarraySum - remainder + p) % p
		if idx, found := lastOccurrence[neededRemainder]; found {
			minLength = min(minLength, i-idx)
		}
		lastOccurrence[subarraySum] = i
	}

	if minLength == len(nums) {
		return -1 
	}

	return minLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{3, 1, 4, 2}
	p := 6
	result := minSubarray(nums, p)
	fmt.Println("The minimum length of the subarray to remove:", result)
	nums2 := []int{6, 3, 5, 2}
	p2 := 9
	result2 := minSubarray(nums2, p2)
	fmt.Println("The minimum length of the subarray to remove:", result2)
}
