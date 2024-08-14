package main

import (
    "fmt"
    "strings"
)

func wordPattern(pattern string, s string) bool {
    patternMap := make(map[string]string)
    wordMap := make(map[string]string)
    words := strings.Split(s, " ")
    
    if len(pattern) != len(words) {
        return false
    }
    
    for i := 0; i < len(pattern); i++ {
        if word, ok := patternMap[string(pattern[i])]; ok {
            if word != words[i] {
                return false
            }
        } else {
            patternMap[string(pattern[i])] = words[i]
        }
        
        if patternChar, ok := wordMap[words[i]]; ok {
            if string(pattern[i]) != patternChar {
                return false
            }
        } else {
            wordMap[words[i]] = string(pattern[i])
        }
    }
    
    return true
}

func main() {
    pattern := "abba"
    s := "dog cat cat dog"
    
    result := wordPattern(pattern, s)
    
    if result {
        fmt.Println("The string follows the pattern.")
    } else {
        fmt.Println("The string does not follow the pattern.")
    }
}
