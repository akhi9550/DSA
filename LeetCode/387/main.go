package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	str := []rune(s)
	mp := make(map[rune]int)
	for _, val := range str {
		_, ok := mp[val]
		if ok {
			mp[val]++
		} else {
			mp[val] = 1
		}
	}
	for i, val := range str {
		pos := mp[val]
		if pos == 1 {
			return i
		}
	}
	return -1
}

func main() {
	input := "loveleetcode"
	result := firstUniqChar(input)
	fmt.Println(result)
}
