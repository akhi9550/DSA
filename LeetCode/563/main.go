package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	res := 0
	postOrder(root, &res)
	return res
}

func postOrder(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	
	left, right := postOrder(root.Left, res), postOrder(root.Right, res)
	*res += abs(left - right)
	
	return left + right + root.Val
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	result := findTilt(root)
	fmt.Println("Tilt of the tree:", result)
}
