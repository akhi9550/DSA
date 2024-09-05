package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	result := []string{}
	helper(root, strconv.Itoa(root.Val), &result)
	return result
}

func helper(root *TreeNode, str string, result *[]string) {
	if root.Left == nil && root.Right == nil {
		*result = append(*result, str)
		return
	}
	if root.Left != nil {
		helper(root.Left, str+"->"+strconv.Itoa(root.Left.Val), result)
	}
	if root.Right != nil {
		helper(root.Right, str+"->"+strconv.Itoa(root.Right.Val), result)
	}
}

func createSampleTree() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	return root
}

func main() {
	root := createSampleTree()
	paths := binaryTreePaths(root)
	fmt.Println("Root-to-leaf paths:", paths)
}
