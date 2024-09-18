package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}
	sort.Slice(strNums, func(i, j int) bool {
		return strNums[i]+strNums[j] > strNums[j]+strNums[i]
	})

	result := ""
	for _, str := range strNums {
		result += str
	}
	if result[0] == '0' {
		return "0"
	}

	return result
}

func main() {
	nums := []int{3, 30, 34, 5, 9}
	result := largestNumber(nums)
	fmt.Println("The largest number is:", result)
}
