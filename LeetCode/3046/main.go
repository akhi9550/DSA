package main

import (
    "fmt"
)

func isPossibleToSplit(nums []int) bool {
    m := map[int]int{}
    for _, num := range nums {
        m[num]++
        if m[num] > 2 {
            return false
        }
    }
    return true
}

func main() {
    nums := []int{1, 2, 2, 3, 3}
    if isPossibleToSplit(nums) {
        fmt.Println("It is possible to split the array.")
    } else {
        fmt.Println("It is not possible to split the array.")
    }
}
