package main

import "fmt"

// TreeNode represents a node in the binary search tree.
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// BuildBSTFromPreorder constructs a Binary Search Tree from a preorder sequence.
func BuildBSTFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Value: preorder[0]}
	i := 1

	// Find the index where the right subtree begins
	for i < len(preorder) && preorder[i] < root.Value {
		i++
	}

	// Recursive construction
	root.Left = BuildBSTFromPreorder(preorder[1:i])
	root.Right = BuildBSTFromPreorder(preorder[i:])

	return root
}

// InorderTraversal performs an in-order traversal and prints the values of the binary search tree.
func InorderTraversal(root *TreeNode) {
	if root != nil {
		InorderTraversal(root.Left)
		fmt.Printf("%d ", root.Value)
		InorderTraversal(root.Right)
	}
}

func main() {
	// Example preorder sequence
	preorder := []int{8, 5, 1, 7, 10, 12}

	// Build BST from the preorder sequence
	root := BuildBSTFromPreorder(preorder)

	// Display the values of the constructed BST using in-order traversal
	fmt.Println("In-order traversal of the constructed BST:")
	InorderTraversal(root)
	fmt.Println()
}
