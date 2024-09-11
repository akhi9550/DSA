package main

import (
    "fmt"
)

func minBitFlips(start int, goal int) int {
    ans := 0
    for start > 0 || goal > 0 {
        if start&1 != goal&1 {
            ans++
        }
        start = start >> 1
        goal = goal >> 1
    }
    return ans
}

func main() {
    var start, goal int
    fmt.Println("Enter the start integer:")
    fmt.Scan(&start)
    fmt.Println("Enter the goal integer:")
    fmt.Scan(&goal)

    flips := minBitFlips(start, goal)
    fmt.Printf("Minimum bit flips to convert %d to %d: %d\n", start, goal, flips)
}
