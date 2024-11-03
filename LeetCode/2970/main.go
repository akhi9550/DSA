package main

import (
    "fmt"
)

func isIncreasing(nums []int, i, j int) bool {
    prev := 0
    for k, num := range nums {
        if k < i || k > j {
            if prev < num {
                prev = num
            } else {
                return false
            }
        }
    }
    return true
}

func incremovableSubarrayCount(nums []int) int {
    result := 0
    for i := range nums {
        for j := i; j < len(nums); j++ {
            if isIncreasing(nums, i, j) {
                result++
            }
        }
    }
    return result
}

func main() {
    nums := []int{1, 3, 2, 4, 5}
    count := incremovableSubarrayCount(nums)
    fmt.Printf("Number of increasing subarrays: %d\n", count)
}
