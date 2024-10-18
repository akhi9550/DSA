package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	return dfs(root, nil)
}

func dfs(root, tail *TreeNode) *TreeNode {
	if root == nil {
		return tail
	}

	res := dfs(root.Left, root)
	root.Left = nil
	root.Right = dfs(root.Right, tail)

	return res
}

func inOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left)
	fmt.Print(root.Val, " ")
	inOrderTraversal(root.Right)
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 8}
	root.Right.Right.Left = &TreeNode{Val: 7}
	root.Right.Right.Right = &TreeNode{Val: 9}
	fmt.Println("Original In-Order Traversal:")
	inOrderTraversal(root)
	fmt.Println()
	newRoot := increasingBST(root)
	fmt.Println("Transformed Tree In-Order Traversal (Increasing Order):")
	inOrderTraversal(newRoot)
	fmt.Println()
}
