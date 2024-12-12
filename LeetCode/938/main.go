package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	return binarySearch(root, low, high)
}

func binarySearch(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}
	if root.Val < low {
		return binarySearch(root.Right, low, high)
	} else if root.Val > high {
		return binarySearch(root.Left, low, high)
	} else {
		return root.Val + binarySearch(root.Left, low, high) + binarySearch(root.Right, low, high)
	}
}

func insertNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insertNode(root.Left, val)
	} else {
		root.Right = insertNode(root.Right, val)
	}
	return root
}

func main() {
	var root *TreeNode
	nums := []int{10, 5, 15, 3, 7, 18}
	for _, num := range nums {
		root = insertNode(root, num)
	}

	low := 7
	high := 15
	result := rangeSumBST(root, low, high)
	fmt.Printf("The sum of values between %d and %d is: %d\n", low, high, result)
}
