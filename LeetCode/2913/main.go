package main

import (
    "fmt"
)

func sumCounts(nums []int) int {
    res := 0
    for i := range nums {
        var n [101]bool
        k := 0
        for _, j := range nums[i:] {
            if !n[j] {
                n[j] = true
                k++
            }
            res += k * k
        }
    }
    return res
}

func main() {
    nums := []int{1, 2, 1, 2, 3}
    result := sumCounts(nums)
    fmt.Println("Result:", result)
}
