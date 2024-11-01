package main

import (
    "fmt"
)

func numberOfAlternatingGroups(colors []int) int {
    colors = append(colors, colors[0], colors[1])

    ans := 0
    for i := 0; i <= len(colors)-3; i++ {
        if colors[i] == 0 && colors[i+1] == 1 && colors[i+2] == 0 {
            ans++
        }
        if colors[i] == 1 && colors[i+1] == 0 && colors[i+2] == 1 {
            ans++
        }
    }
    return ans
}

func main() {
    colors := []int{0, 1, 0, 1, 0}
    result := numberOfAlternatingGroups(colors)
    fmt.Println("Number of alternating groups:", result)
}
