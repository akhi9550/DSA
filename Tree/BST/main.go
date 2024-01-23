package main

import (
	"fmt"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (n *TreeNode) Insert(value int) *TreeNode {
	if n == nil {
		return &TreeNode{Value: value}
	}

	if value <= n.Value {
		n.Left = n.Left.Insert(value)
	} else {
		n.Right = n.Right.Insert(value)
	}

	return n
}

func (n *TreeNode) Search(value int) *TreeNode {
	if n == nil || value == n.Value {
		return n
	}

	if value < n.Value {
		return n.Left.Search(value)
	}
	return n.Right.Search(value)
}

func (n *TreeNode) Delete(value int) *TreeNode {
	if n == nil {
		return nil
	}

	if value < n.Value {
		n.Left = n.Left.Delete(value)
	} else if value > n.Value {
		n.Right = n.Right.Delete(value)
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		// Find the inorderSucessor
		n.Value = findMinValue(n.Right)
		n.Right = n.Right.Delete(n.Value)
	}
	return n
}
func findMinValue(n *TreeNode) int {
	for n.Left != nil {
		n = n.Left
	}
	return n.Value
}

func (n *TreeNode) Contains(value int) bool {
	return n.Search(value) != nil
}

func (n *TreeNode) Display() {
	if n != nil {
		n.Left.Display()
		fmt.Printf("%d ", n.Value)
		n.Right.Display()
	}
}
func main() {
	root := &TreeNode{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	fmt.Println("In-order traversal:")
	root.Display()
	fmt.Println()

	fmt.Println("Contains 7:", root.Contains(7))
	fmt.Println("Contains 11:", root.Contains(11))

	fmt.Println("min", findMinValue(root))

	root = root.Delete(7)
	fmt.Println("After deleting 7:")
	root.Display()

}
