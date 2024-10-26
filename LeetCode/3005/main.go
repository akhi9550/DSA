package main

import (
    "fmt"
)

func maxFrequencyElements(nums []int) int {
    var f [101]int
    maxFreq := 0
    res := 0

    for _, n := range nums {
        f[n]++
        
        if f[n] > maxFreq {
            maxFreq = f[n]
            res = n
        } else if f[n] == maxFreq {
            res += n
        }
    }
    return res
}

func main() {
    nums := []int{1, 3, 2, 3, 4, 2, 3}
    result := maxFrequencyElements(nums)
    fmt.Printf("The sum of elements with the maximum frequency is: %d\n", result)
}
