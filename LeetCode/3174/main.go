package main

import (
	"fmt"
)

func clearDigits(s string) string {
	for {
		var found bool
		for i, l := range s {
			if l > 47 && l < 58 {
				if i != 0 && (s[i-1] < 48 || s[i-1] > 57) {
					if i == len(s)-1 {
						s = string(s[:i-1])
					} else {
						s = string(s[:i-1]) + string(s[i+1:])
					}
				}
				found = true
				break
			}
		}
		if !found {
			break
		}
	}
	return s
}

func main() {
	input := "a1b2c3d"
	result := clearDigits(input)
	fmt.Println("Original String:", input)
	fmt.Println("Processed String:", result)
}
