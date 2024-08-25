package main

import (
    "fmt"
)

func combinationSum4(nums []int, target int) int {
    dict := make(map[int]int)
    return helper(nums, target, dict)
}

func helper(nums []int, target int, dict map[int]int) int {
    if v, ok := dict[target]; ok {
        return v
    }
    
    if target == 0 {
        return 1
    }
    
    if target < 0 {
        return 0
    }
    
    var res int
    for _, v := range nums {
        res += helper(nums, target-v, dict)
    }
    
    dict[target] = res
    
    return res
}

func main() {
    nums := []int{1, 2, 3}
    target := 4
    result := combinationSum4(nums, target)
    fmt.Printf("Number of combinations for target %d with nums %v: %d\n", target, nums, result)
}
