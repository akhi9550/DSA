package main

import (
	"fmt"
)

func countQuadruplets(nums []int) int {
	counter := 0
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for k := j + 1; k < len(nums)-1; k++ {
				for l := k + 1; l < len(nums); l++ {
					if nums[i]+nums[j]+nums[k] == nums[l] {
						counter++
					}
				}
			}
		}
	}
	return counter
}

func main() {
	nums := []int{1, 2, 3, 6} 
	result := countQuadruplets(nums)
	fmt.Println("Number of quadruplets:", result)
}
