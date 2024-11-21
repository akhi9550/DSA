package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var result int
	maxLength(root, &result)
	return result
}

func maxLength(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	lenLeft := maxLength(root.Left, result)
	lenRight := maxLength(root.Right, result)
	*result = max(lenLeft+lenRight, *result)

	return 1 + max(lenLeft, lenRight)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println("Diameter of the binary tree:", diameterOfBinaryTree(root)) // Output: 3
}
