package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	if len(words) == 0 {
		return 0
	}
	return len(words[len(words)-1])
}

func main() {
	s := "Hello World"
	length := lengthOfLastWord(s)
	fmt.Printf("Length of the last word in '%s' is: %d\n", s, length)
}
