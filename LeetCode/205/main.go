package main

import (
    "fmt"
)

func isIsomorphic(s string, t string) bool {
    ms := make(map[uint8]uint8)
    mt := make(map[uint8]uint8)

    for i := 0; i < len(s); i++ {
        if val, ok := ms[s[i]]; ok && val != t[i] {
            return false
        }

        if val, ok := mt[t[i]]; ok && val != s[i] {
            return false
        }

        ms[s[i]] = t[i]
        mt[t[i]] = s[i]
    }

    return true
}

func main() {
    s1 := "egg"
    t1 := "add"
    s2 := "foo"
    t2 := "bar"
    s3 := "paper"
    t3 := "title"
    s4 := "ab"
    t4 := "aa"

    fmt.Printf("Is '%s' isomorphic to '%s'? %v\n", s1, t1, isIsomorphic(s1, t1)) // true
    fmt.Printf("Is '%s' isomorphic to '%s'? %v\n", s2, t2, isIsomorphic(s2, t2)) // false
    fmt.Printf("Is '%s' isomorphic to '%s'? %v\n", s3, t3, isIsomorphic(s3, t3)) // true
    fmt.Printf("Is '%s' isomorphic to '%s'? %v\n", s4, t4, isIsomorphic(s4, t4)) // false
}
