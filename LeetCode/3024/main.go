package main

import (
    "fmt"
    "sort"
)

func triangleType(nums []int) string {
    sort.Ints(nums)
    if nums[0]+nums[1] <= nums[2] {
        return "none"
    }
    if nums[0] == nums[1] && nums[1] == nums[2] {
        return "equilateral"
    }
    if nums[0] == nums[1] || nums[1] == nums[2] {
        return "isosceles"
    }
    return "scalene"
}

func main() {
    triangles := [][]int{
        {3, 4, 5},   
        {2, 2, 3},   
        {2, 2, 2},   
        {1, 2, 3},    
    }

    for _, sides := range triangles {
        fmt.Printf("Triangle with sides %v is %s\n", sides, triangleType(sides))
    }
}
