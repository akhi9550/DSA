package main

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

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
			if current.Right != nil {
				successor = findMin(current.Right)
			}
			break
		}
	}

	return successor
}

func findMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func main() {
	root := &TreeNode{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	fmt.Println("In-order traversal of the original BST:")
	printInorder(root)
	fmt.Println()

	key := 7
	successor := InorderSuccessor(root, key)

	if successor != nil {
		fmt.Printf("Inorder successor for key %d is: %d\n", key, successor.Value)
	} else {
		fmt.Printf("No inorder successor found for key %d\n", key)
	}
}

func printInorder(root *TreeNode) {
	if root != nil {
		printInorder(root.Left)
		fmt.Printf("%d ", root.Value)
		printInorder(root.Right)
	}
}
