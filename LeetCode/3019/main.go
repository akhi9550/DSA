package main

import (
    "fmt"
    "strings"
)

func countKeyChanges(s string) int {
    var count int
    c := strings.ToLower(s)
    for i := 1; i < len(s); i++ {
        if c[i] != c[i-1] {
            count++
        }
    }
    return count
}

func main() {
    input := "aAaBBbCcC"
    result := countKeyChanges(input)
    fmt.Printf("The number of key changes in the string is: %d\n", result)
}
