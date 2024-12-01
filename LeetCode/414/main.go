package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	firstMax, secondMax, thirdMax := math.MinInt64, math.MinInt64, math.MinInt64

	for _, num := range nums {
		if num > firstMax {
			thirdMax = secondMax
			secondMax = firstMax
			firstMax = num
		} else if num < firstMax && num > secondMax {
			thirdMax = secondMax
			secondMax = num
		} else if num < secondMax && num > thirdMax {
			thirdMax = num
		}
	}

	if thirdMax != math.MinInt64 {
		return thirdMax
	}

	return firstMax
}

func main() {
	fmt.Println(thirdMax([]int{3, 2, 1}))       // Output: 1
	fmt.Println(thirdMax([]int{1, 2}))          // Output: 2
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))    // Output: 1
	fmt.Println(thirdMax([]int{5, 5, 5, 5}))    // Output: 5
	fmt.Println(thirdMax([]int{10, 20, 30, 40})) // Output: 20
}
