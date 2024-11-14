package main

import (
    "fmt"
    "sort"
)

func maximumProduct(nums []int) int {
    sort.Ints(nums)
    return max(nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3], nums[0]*nums[1]*nums[len(nums)-1])
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    nums := []int{-10, -10, 5, 2}
    result := maximumProduct(nums)
    fmt.Println("The maximum product of three numbers is:", result)
}
