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
func main() {
	root := &TreeNode{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)
	fmt.Println()

	target := 8
	closestValue := root.FindClosestValue(target)
	fmt.Printf("Closest value to %d is: %d\n", target, closestValue)
}
