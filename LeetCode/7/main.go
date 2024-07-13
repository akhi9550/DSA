package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	negative := false
	if x < 0 {
		negative = true
		x = -x
	}
	str := strconv.Itoa(x)
	v := []rune(str)
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}
	res, _ := strconv.Atoi(string(v))
	if negative {
		res = -res
	}
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}

	return res
}

func main() {
	fmt.Println(reverse(123))  // Output: 321
	fmt.Println(reverse(-123)) // Output: -321
	fmt.Println(reverse(120))  // Output: 21
}
