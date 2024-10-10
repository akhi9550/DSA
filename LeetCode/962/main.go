package main

import (
    "fmt"
    "math"
    "sort"
)

func maxWidthRamp(A []int) int {
    b := make([][2]int, len(A))
    for i := 0; i < len(b); i++ {
        b[i] = [2]int{A[i], i}
    }
    sort.Slice(b, func(i int, j int) bool {
        if b[i][0] == b[j][0] {
            return b[i][1] < b[j][1]
        }
        return b[i][0] < b[j][0]
    })
    
    min := b[0][1]
    res := math.MinInt64
    for i := 1; i < len(b); i++ {
        if b[i][1] - min > res {
            res = b[i][1] - min
        }
        if b[i][1] < min {
            min = b[i][1]
        }
    }
    if res < 0 {
        return 0
    }
    return res
}

func main() {
    A := []int{6, 0, 8, 2, 1, 5}
    result := maxWidthRamp(A)
    fmt.Println("Max width ramp:", result)
}
