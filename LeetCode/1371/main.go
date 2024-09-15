package main

import (
	"fmt"
)

func findTheLongestSubstring(s string) int {
	counter := make([][5]int, len(s)+1)
	a, e, i, o, u := 0, 0, 0, 0, 0

	for index, v := range s {
		switch v {
		case 'a':
			a++
		case 'e':
			e++
		case 'i':
			i++
		case 'o':
			o++
		case 'u':
			u++
		}
		counter[index+1][0] = a
		counter[index+1][1] = e
		counter[index+1][2] = i
		counter[index+1][3] = o
		counter[index+1][4] = u
	}

	for wide := len(s); wide > 0; wide-- {
		for i := 0; i+wide <= len(s); i++ {
			if (counter[i+wide][0]-counter[i][0])%2 == 1 {
				continue
			}
			if (counter[i+wide][1]-counter[i][1])%2 == 1 {
				continue
			}
			if (counter[i+wide][2]-counter[i][2])%2 == 1 {
				continue
			}
			if (counter[i+wide][3]-counter[i][3])%2 == 1 {
				continue
			}
			if (counter[i+wide][4]-counter[i][4])%2 == 1 {
				continue
			}
			return wide
		}
	}
	return 0
}

func main() {
	s := "eleetminicoworoep"
	fmt.Println("Test case 1:", findTheLongestSubstring(s)) // Expected output: 13
	s2 := "leetcodeisgreat"
	fmt.Println("Test case 2:", findTheLongestSubstring(s2)) // Expected output: 5
	s3 := "bcbcbc"
	fmt.Println("Test case 3:", findTheLongestSubstring(s3)) // Expected output: 6
}
