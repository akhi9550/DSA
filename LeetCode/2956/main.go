package main

import (
	"fmt"
)

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]struct{})
	map2 := make(map[int]struct{})
	for _, v := range nums1 {
		map1[v] = struct{}{}
	}

	for _, v := range nums2 {
		map2[v] = struct{}{}
	}

	result := make([]int, 2)

	for _, val := range nums1 {
		if _, ok := map2[val]; ok {
			result[0]++
		}
	}

	for _, val := range nums2 {
		if _, ok := map1[val]; ok {
			result[1]++
		}
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{4, 5, 6, 7, 8}
	result := findIntersectionValues(nums1, nums2)
	fmt.Println("Count of elements in nums1 that intersect with nums2:", result[0])
	fmt.Println("Count of elements in nums2 that intersect with nums1:", result[1])
}
