package main

import (
    "fmt"
)

func convertToTitle(n int) string {
    result := ""
    for {
        if n <= 26 {
            result = string(n+64) + result
            break
        } else {
            result = string((n-1)%26 + 65) + result
            n = (n-1) / 26
        }
    }
    return result
}

func main() {
    testCases := []int{1, 26, 27, 52, 701, 702}
    for _, n := range testCases {
        fmt.Printf("Input: %d, Output: %s\n", n, convertToTitle(n))
    }
}
