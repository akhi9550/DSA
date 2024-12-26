package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	sign := 1
	i := 0
	if s[0] == '+' || s[0] == '-' {
		if s[0] == '-' {
			sign = -1
		}
		i++
	}

	result := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		result = result*10 + int(s[i]-'0')
		if sign*result > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign*result < math.MinInt32 {
			return math.MinInt32
		}
		i++
	}

	return sign * result
}

func main() {
	fmt.Println(myAtoi("42"))              // Output: 42
	fmt.Println(myAtoi("   -42"))          // Output: -42
	fmt.Println(myAtoi("4193 with words")) // Output: 4193
	fmt.Println(myAtoi("words and 987"))   // Output: 0
	fmt.Println(myAtoi("-91283472332"))    // Output: -2147483648
}
