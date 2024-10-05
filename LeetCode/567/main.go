package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	var pattern, arr [26]int
	for i := 0; i < len(s1); i++ {
		pattern[int(s1[i]-'a')]++
		arr[int(s2[i]-'a')]++
	}

	i := 0
	for {
		if arr == pattern {
			return true
		}
		if i+len(s1) >= len(s2) {
			break
		}
		arr[int(s2[i]-'a')]--
		arr[int(s2[i+len(s1)]-'a')]++
		i++
	}
	return false
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"

	if checkInclusion(s1, s2) {
		fmt.Println("s1's permutation is a substring of s2.")
	} else {
		fmt.Println("s1's permutation is not a substring of s2.")
	}
}
