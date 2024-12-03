package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	return findMinNode(root, root.Val)
}

func findMinNode(root *TreeNode, except int) int {
	if root == nil {
		return -1
	}

	if root.Val != except {
		return root.Val
	}

	l := findMinNode(root.Left, except)
	r := findMinNode(root.Right, except)

	if l == -1 || (r != -1 && r < l) {
		return r
	}
	return l
}

func createTree(vals []int, index int) *TreeNode {
	if index >= len(vals) || vals[index] == -1 {
		return nil
	}

	node := &TreeNode{Val: vals[index]}
	node.Left = createTree(vals, 2*index+1)
	node.Right = createTree(vals, 2*index+2)
	return node
}

func main() {
	tests := []struct {
		values []int
		output int
	}{
		{[]int{2, 2, 5, -1, -1, 5, 7}, 5},
		{[]int{2, 2, 2}, -1},
		{[]int{2, 2, 3}, 3},
		{[]int{5}, -1},
		{[]int{1, 1, 1, 1, 1, 1, 2}, 2},
	}

	for i, test := range tests {
		root := createTree(test.values, 0)
		result := findSecondMinimumValue(root)
		fmt.Printf("Test %d: Expected %d, Got %d\n", i+1, test.output, result)
	}
}
