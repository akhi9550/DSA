package main

import (
	"fmt"
)

func shortestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	
	tmp := make([]byte, len(s), 2*len(s))
	i := 0
	for i < len(s) {
		tmp[i] = s[len(s)-i-1]
		i++
	}

	tmps := s + "#" + string(tmp)
	next := make([]int, len(tmps))
	k := -1
	next[0] = -1
	i = 0
	
	for i < len(tmps)-1 {
		if k == -1 || tmps[i] == tmps[k] {
			k++
			i++
			next[i] = k
		} else {
			k = next[k]
		}
	}
	
	leng := next[len(next)-1] + 1
	if leng > len(s) {
		return s
	}
	
	rst := tmps[len(s)+1:len(tmps)-leng] + s
	return rst
}

func main() {
	str1 := "aacecaaa"
	str2 := "abcd"

	fmt.Println("Shortest Palindrome of", str1, "is:", shortestPalindrome(str1))
	fmt.Println("Shortest Palindrome of", str2, "is:", shortestPalindrome(str2))
}
