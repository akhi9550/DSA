package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	a := 0
	for i := 0; i < len(t); i++ {
		if a < len(s) && s[a] == t[i] {
			a++
		}
	}
	return len(s) == a
}

func main() {
	s1 := "abc"
	t1 := "ahbgdc"
	fmt.Printf("Is '%s' a subsequence of '%s'? %v\n", s1, t1, isSubsequence(s1, t1))

	s2 := "axc"
	t2 := "ahbgdc"
	fmt.Printf("Is '%s' a subsequence of '%s'? %v\n", s2, t2, isSubsequence(s2, t2))

	s3 := "hello"
	t3 := "heyllo world"
	fmt.Printf("Is '%s' a subsequence of '%s'? %v\n", s3, t3, isSubsequence(s3, t3))
}
