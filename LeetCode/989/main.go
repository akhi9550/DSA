package main

import "fmt"

func addToArrayForm(num []int, k int) []int {
    sum := k
    ans := make([]int, 0)
    for i := len(num) - 1; i >= 0 || sum > 0; i-- {
        if i >= 0 {
            sum += num[i]
        }
        ans = append(ans, sum%10)
        sum /= 10
    }
    for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
        ans[i], ans[j] = ans[j], ans[i]
    }
    return ans
}

func main() {
    num := []int{1, 2, 0, 0}
    k := 34
    result := addToArrayForm(num, k)
    fmt.Println(result)
}
