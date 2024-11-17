package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
    total := 0
    last := 0
    seconds := 0

    for _, val := range timeSeries {
        if seconds > 0 {
            if val > seconds {
                total += last
            } else {
                left := seconds - val + 1
                total += last - left
            }
            last = duration
            seconds = val + duration - 1
        } else {
            total += last
            seconds = val + duration - 1
            last = duration
        }
    }
    return total + last
}

func main() {
    timeSeries := []int{1, 4, 5}
    duration := 2

    result := findPoisonedDuration(timeSeries, duration)
    fmt.Printf("Total poisoned duration: %d\n", result)
}
