package main

import (
    "fmt"
)

type NumArray struct {
    sum []int
}

func Constructor(nums []int) NumArray {
    if len(nums) == 0 {
        return NumArray{}
    }
    
    sum := make([]int, len(nums) + 1)
    sum[0], sum[1] = 0, nums[0]
    for i := 2; i < len(sum); i++ {
        sum[i] = sum[i - 1] + nums[i - 1]
    }
    return NumArray{sum: sum}
}

func (this *NumArray) SumRange(i int, j int) int {
    return this.sum[j + 1] - this.sum[i]
}

func main() {
    nums := []int{-2, 0, 3, -5, 2, -1}
    
    numArray := Constructor(nums)
    
    sum1 := numArray.SumRange(0, 2)
    fmt.Println("Sum of elements from index 0 to 2:", sum1) // Expected output: 1

    sum2 := numArray.SumRange(2, 5)
    fmt.Println("Sum of elements from index 2 to 5:", sum2) // Expected output: -1

    sum3 := numArray.SumRange(0, 5)
    fmt.Println("Sum of elements from index 0 to 5:", sum3) // Expected output: -3
}
