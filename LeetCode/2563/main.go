package main

import (
    "fmt"
    "sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
    result := int64(0)
    sort.Ints(nums)
    
    for i := 0; i < len(nums); i++ {
        start := lowerBound(nums, i+1, len(nums)-1, lower-nums[i])
        end := upperBound(nums, i+1, len(nums)-1, upper-nums[i])
        result += int64(end - start)
    }
    
    return result
}

func lowerBound(nums []int, left, right, target int) int {
    for left <= right {
        mid := (left + right) / 2
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return left
}

func upperBound(nums []int, left, right, target int) int {
    for left <= right {
        mid := (left + right) / 2
        if nums[mid] <= target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return left
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    lower := 4
    upper := 6

    result := countFairPairs(nums, lower, upper)
    fmt.Printf("The count of fair pairs is: %d\n", result)
}
