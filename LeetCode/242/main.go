package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	dict := make(map[rune]int)
	for _, v := range s {
		dict[v]++
	}
	for _, v := range t {
		dict[v]--
	}
	for _, v := range dict {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "listen"
	t := "silent"
	if isAnagram(s, t) {
		fmt.Printf("'%s' and '%s' are anagrams.\n", s, t)
	} else {
		fmt.Printf("'%s' and '%s' are not anagrams.\n", s, t)
	}

	s = "hello"
	t = "world"
	if isAnagram(s, t) {
		fmt.Printf("'%s' and '%s' are anagrams.\n", s, t)
	} else {
		fmt.Printf("'%s' and '%s' are not anagrams.\n", s, t)
	}
}
