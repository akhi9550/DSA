package main

import (
	"fmt"
	"math"
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

		// Find the minimum value in the right subtree
		min := n.Right.findMinValue()
		n.Value = min
		// Delete the node with the minimum value in the right subtree
		n.Right = n.Right.Delete(min)
	}
	return n
}

func (n *TreeNode) FindClosestValue(target int) int {
	closest := n.Value
	current := n

	for current != nil {
		if abs(target-closest) > abs(target-current.Value) {
			closest = current.Value
		}

		if target < current.Value {
			current = current.Left
		} else if target > current.Value {
			current = current.Right
		} else {
			break
		}
	}

	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (n *TreeNode) findMinValue() int {
	for n.Left != nil {
		n = n.Left
	}
	return n.Value
}

func (n *TreeNode) Validate() bool {
	return n.validateHelper(math.MinInt, math.MaxInt)
}

func (n *TreeNode) validateHelper(min, max int) bool {
	if n == nil {
		return true
	}

	if n.Value <= min || n.Value >= max {
		return false
	}

	return n.Left.validateHelper(min, n.Value) && n.Right.validateHelper(n.Value, max)
}

// Contains checks if a value exists in the binary search tree.
func (n *TreeNode) Contains(value int) bool {
	return n.Search(value) != nil
}

// Display in-order traversal
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

	fmt.Println("min", root.findMinValue())

	root = root.Delete(7)
	fmt.Println("After deleting 7:")
	root.Display()
	fmt.Println()

	target := 8
	closestValue := root.FindClosestValue(target)
	fmt.Printf("Closest value to %d is: %d\n", target, closestValue)

	fmt.Println("Is the tree a BST :-", root.Validate())
}
