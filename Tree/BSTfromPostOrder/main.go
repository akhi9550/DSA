package main

import "fmt"

// TreeNode represents a node in the binary search tree.
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// BuildBSTFromPostorder constructs a Binary Search Tree from a postorder sequence.
func BuildBSTFromPostorder(postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Value: postorder[len(postorder)-1]}
	i := 0

	// Find the index where the left subtree ends
	for i < len(postorder)-1 && postorder[i] < root.Value {
		i++
	}

	// Recursive construction
	root.Left = BuildBSTFromPostorder(postorder[:i])
	root.Right = BuildBSTFromPostorder(postorder[i : len(postorder)-1])

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
	// Example postorder sequence
	postorder := []int{1, 7, 5, 12, 10, 8}

	// Build BST from the postorder sequence
	root := BuildBSTFromPostorder(postorder)

	// Display the values of the constructed BST using in-order traversal
	fmt.Println("In-order traversal of the constructed BST:")
	InorderTraversal(root)
	fmt.Println()
}
