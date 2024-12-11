package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func commonChars(a []string) []string {
	res := make([]int, 26)

	for i := range res {
		res[i] = math.MaxInt32
	}

	for _, s := range a {
		count := make([]int, 26)
		for _, char := range s {
			count[char-'a']++
		}
		for i := 0; i < 26; i++ {
			res[i] = min(res[i], count[i])
		}
	}

	var resArray []string
	for i := 0; i < 26; i++ {
		for res[i] != 0 {
			resArray = append(resArray, string('a'+i))
			res[i]--
		}
	}
	return resArray
}

func main() {
	input := []string{"bella", "label", "roller"}
	result := commonChars(input)
	fmt.Println("Common characters:", result)
}