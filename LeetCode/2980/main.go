package main

import "fmt"

func hasTrailingZeros(nums []int) bool {
    even := 0
    for i := range nums {
        if nums[i]%2 == 0 {
            even++
            if even == 2 {
                return true
            }
        }
    }
    return false
}

func main() {
    nums := []int{1, 3, 5, 8, 10}
    fmt.Println(hasTrailingZeros(nums))
}
