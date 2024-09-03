package main

import (
    "fmt"
)

func getLucky(s string, k int) int {
    res := 0
    
    for _, v := range s {
        n := int(v - 'a') + 1
        
        res += (n % 10) + (n / 10)
    }
    
    for i := 0; i < k-1; i++ {
        res = transform(res)
    }
    
    return res
}

func transform(n int) int {
    res := 0
        for n != 0 {
        res += n % 10
        n /= 10
    }
    
    return res
}

func main() {
    s := "helloworld"
    k := 2
        result := getLucky(s, k)
    fmt.Println("The result is:", result)
}
