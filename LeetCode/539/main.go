package main

import (
	"fmt"
	"sort"
	"strconv"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMinDifference(timePoints []string) int {
	minutes := []int{}
	for _, tp := range timePoints {
		a, _ := strconv.Atoi(tp[:2])
		b, _ := strconv.Atoi(tp[3:])
		minutes = append(minutes, a*60+b)
	}

	sort.Ints(minutes)

	res := 25 * 60
	for i := 1; i < len(minutes); i++ {
		res = min(res, minutes[i]-minutes[i-1])
	}

	res = min(res, 24*60-(minutes[len(minutes)-1]-minutes[0]))

	return res
}

func main() {
	timePoints := []string{"23:59", "00:00", "05:30", "12:15"}
	fmt.Println("Minimum time difference is:", findMinDifference(timePoints))
}
