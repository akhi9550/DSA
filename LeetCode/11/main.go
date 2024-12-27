package main

import "fmt"

func maxArea(height []int) int {
    var max int
    for i, j := 0, len(height)-1; i < j; {
        width := j - i
        sum := width * min(height[i], height[j])
        if sum > max {
            max = sum
        }
        if height[i] < height[j] {
            i++
        } else {
            j--
        }
    }
    return max
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
    fmt.Println("Maximum Area:", maxArea(heights)) // Output: 49
}
