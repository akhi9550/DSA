package main

import (
    "fmt"
)

func findPeaks(mountain []int) []int {
    var peaks []int
    j := 0
    for i := 1; i < len(mountain)-1; i++ {
        if mountain[j] < mountain[i] && mountain[i] > mountain[i+1] {
            peaks = append(peaks, i)
        }
        j++
    }
    return peaks
}

func main() {
    mountain := []int{1, 3, 7, 6, 5, 10, 6, 3, 2}
    peaks := findPeaks(mountain)
    fmt.Println("Peaks are at indices:", peaks)
}
