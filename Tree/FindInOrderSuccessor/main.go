package main

import "fmt"

// TreeNode represents a node in the binary search tree.
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// Insert inserts a value into the binary search tree.
func (n *TreeNode) Insert(value int) *TreeNode {
	if n == nil {
		return &TreeNode{Value: value}
	}

	if value < n.Value {
		n.Left = n.Left.Insert(value)
	} else if value > n.Value {
		n.Right = n.Right.Insert(value)
	}

	return n
}

// InorderSuccessor finds the inorder successor for a given key in the binary search tree.
func InorderSuccessor(root *TreeNode, key int) *TreeNode {
	var successor *TreeNode
	current := root

	for current != nil {
		if key < current.Value {
			successor = current
			current = current.Left
		} else if key > current.Value {
			current = current.Right
		} else {
			// If the key is found, check if there is a right subtree
			if current.Right != nil {
				// Find the minimum value in the right subtree
				successor = findMin(current.Right)
			}
			break
		}
	}

	return successor
}

// findMin finds the node with the minimum key in the binary search tree.
func findMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func main() {
	// Creating a sample binary search tree
	root := &TreeNode{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	// Displaying the values of the original binary search tree in order
	fmt.Println("In-order traversal of the original BST:")
	printInorder(root)
	fmt.Println()

	// Find the inorder successor for the key 7
	key := 7
	successor := InorderSuccessor(root, key)

	if successor != nil {
		fmt.Printf("Inorder successor for key %d is: %d\n", key, successor.Value)
	} else {
		fmt.Printf("No inorder successor found for key %d\n", key)
	}
}

// printInorder performs an in-order traversal and prints the values of the binary search tree.
func printInorder(root *TreeNode) {
	if root != nil {
		printInorder(root.Left)
		fmt.Printf("%d ", root.Value)
		printInorder(root.Right)
	}
}
