package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	value       int
	left, right *TreeNode
}

func (t *TreeNode) Insert(data int) *TreeNode {
	if t == nil {
		return &TreeNode{value: data}
	}
	if data <= t.value {
		t.left = t.left.Insert(data)
	} else if data > t.value {
		t.right = t.right.Insert(data)
	}
	return t
}

func (t *TreeNode) Display() {
	if t != nil {
		t.left.Display()
		fmt.Println(" ", t.value)
		t.right.Display()
	}
}

func (t *TreeNode) Validate(min, max int) bool {
	if t == nil {
		return true
	}
	if t.value <= min || t.value >= max {
		return false
	}
	return t.left.Validate(min, t.value) && t.right.Validate(t.value, max)
}

func FindNthSmallest(root *TreeNode, n int) int {
	count := 0
	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		count++
		if count == n {
			return current.value
		}

		current = current.right
	}

	return 0
}

func main() {
	root := &TreeNode{value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)
	root.Display()
	fmt.Println("is bsf", root.Validate(math.MinInt, math.MaxInt))
	n := 4
	value := FindNthSmallest(root, n)
	fmt.Printf("The %d-th smallest value in the BST is: %d\n", n, value)
}
