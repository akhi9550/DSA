package main

import "fmt"

func countLargestGroup(n int) int {
    tmp := make(map[int]int)
    maxSize := 0
    res := 0
    for i := 1; i <= n; i++ {
        temp := helper(i)
        tmp[temp]++
        if tmp[temp] > maxSize {
            maxSize = tmp[temp]
        }
    }
    for _, v := range tmp {
        if v == maxSize {
            res++
        }
    }
    return res
}

func helper(n int) int {
    m := 0
    for n > 0 {
        m += n % 10
        n /= 10
    }
    return m
}

func main() {
    n := 13
    result := countLargestGroup(n)
    fmt.Println(result)
}
