package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	arr := []int{}
	isValid(root, &arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}

func isValid(root *TreeNode, arr *[]int) {
	if root != nil {
		isValid(root.Left, arr)
		*arr = append(*arr, root.Val)
		isValid(root.Right, arr)
	}
}

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	if isValidBST(root) {
		fmt.Println("The tree is a valid BST.")
	} else {
		fmt.Println("The tree is not a valid BST.")
	}
}
