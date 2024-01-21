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
func main() {
	root := &TreeNode{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)
	fmt.Println("Is the tree a BST :-", root.Validate())
}
