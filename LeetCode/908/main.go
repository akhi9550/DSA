package main

import (
    "fmt"
)

func smallestRangeI(nums []int, k int) int {
    least := 0

    if len(nums) == 1 {
        return 0
    } else {
        min := nums[0]
        max := nums[0]

        for _, value := range nums {
            if value < min {
                min = value
            }
            if value > max {
                max = value
            }
        }

        least = (max - k) - (min + k)
        if least <= 0 {
            return 0
        }
    }
    return least
}

func main() {
    nums := []int{1, 3, 6}
    k := 3

    result := smallestRangeI(nums, k)
    fmt.Printf("The smallest possible range is: %d\n", result)
}
