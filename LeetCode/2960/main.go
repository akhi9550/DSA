package main

import "fmt"

func countTestedDevices(batteryPercentages []int) int {
    res := 0
    for _, i := range batteryPercentages {
        if i > res {
            res++
        }
    }
    return res
}

func main() {
    batteryPercentages := []int{20, 40, 60, 80, 100}
    result := countTestedDevices(batteryPercentages)
    fmt.Println(result)
}
