package main

import (
    "fmt"
)

func reverseVowels(s string) string {
    res := make([]byte, len(s))
    
    vow := make(map[byte]bool)
    vow['a'] = true
    vow['e'] = true
    vow['i'] = true
    vow['o'] = true
    vow['u'] = true
    vow['A'] = true
    vow['E'] = true
    vow['I'] = true
    vow['O'] = true
    vow['U'] = true
    
    i, j := 0, len(s)-1
    
    for i <= j {
        if !vow[s[i]] {
            res[i] = s[i]
            i++
            continue
        }
        if !vow[s[j]] {
            res[j] = s[j]
            j--
            continue
        }
        res[i] = s[j]
        res[j] = s[i]
        i++
        j--
    }
    return string(res)
}

func main() {
    input := "hello world"
    output := reverseVowels(input)
    fmt.Println("Input: ", input)
    fmt.Println("Output: ", output)
}
