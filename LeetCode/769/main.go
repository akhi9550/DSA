package main

import "fmt"

func maxChunksToSorted(arr []int) int {
    m, l, rs := len(arr)-1, len(arr)-1, 0
    for j := len(arr) - 1; j >= 0; j-- {
        if arr[j] < m {
            m = arr[j]
            l = arr[j]
        }
        if j == l {
            rs++
        }
    }
    return rs
}

func main() {
    arr := []int{4, 3, 2, 1, 0}
    result := maxChunksToSorted(arr)
    fmt.Println(result)
}
