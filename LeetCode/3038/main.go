package main

import (
    "fmt"
)

func maxOperations(nums []int) int {
    n := len(nums)
    val := nums[0] + nums[1]
    ans := 1
    for i := 3; i < n; i += 2 {
        if nums[i]+nums[i-1] == val {
            ans++
        } else {
            break
        }
    }
    return ans
}

func main() {
    nums := []int{4, 1, 5, 0, 4, 1, 5, 0} 
    result := maxOperations(nums)
    fmt.Println("Max Operations:", result)
}
