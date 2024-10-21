package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	queue := []*TreeNode{root}
	sums := []int{}
	for len(queue) != 0 {
		sum := 0
		for i := len(queue); i > 0; i-- {
			curr := queue[0]
			queue = queue[1:]
			if curr == nil {
				continue
			}
			sum += curr.Val
			queue = append(queue, curr.Left, curr.Right)
		}
		if sum != 0 {
			sums = append(sums, sum)
		}
	}
	n := len(sums)
	if k > n {
		return -1
	}
	sort.Ints(sums)
	return int64(sums[n-k])
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 10}
	k := 2
	result := kthLargestLevelSum(root, k)
	fmt.Printf("The %dth largest level sum is: %d\n", k, result)
}
