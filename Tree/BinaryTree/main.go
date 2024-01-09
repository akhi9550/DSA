package main

import (
	"fmt"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (t *TreeNode) Insert(value int) *TreeNode {
	if t == nil {
		return &TreeNode{Value: value}
	}
	if value <= t.Value {
		t.Left = t.Left.Insert(value)
	} else {
		t.Right = t.Right.Insert(value)
	}
	return t
}
func (t *TreeNode) Search(value int) bool {
	if t == nil {
		return false
	}
	if value == t.Value {
		return true
	} else if value < t.Value {
		return t.Left.Search(value)
	} else {
		return t.Right.Search(value)
	}
}

func (t *TreeNode) Delete(value int) *TreeNode {
	if t == nil {
		return nil
	}

	if value < t.Value {
		t.Left = t.Left.Delete(value)
	} else if value > t.Value {
		t.Right = t.Right.Delete(value)
	} else {
		if t.Left == nil {
			return t.Right
		} else if t.Right == nil {
			return t.Left
		}

		// Node with two children: Get the inorder successor (smallest in the right subtree)
		t.Value = minValue(t.Right)

		// Delete the inorder successor
		t.Right = t.Right.Delete(t.Value)
	}
	return t
}

func minValue(t *TreeNode) int {
	for t.Left != nil {
		t = t.Left
	}
	return t.Value
}

func (t *TreeNode) DisplayBFS() {
	queue := []*TreeNode{t}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func (t *TreeNode) DisplayIn() {
	if t != nil {
		t.Left.DisplayIn()
		fmt.Printf("%d ", t.Value)
		t.Right.DisplayIn()
	}
}

func (t *TreeNode) DisplayPre() {
	if t != nil {
		fmt.Printf("%d ", t.Value)
		t.Left.DisplayPre()
		t.Right.DisplayPre()
	}
}

func (t *TreeNode) DisplayPost() {
	if t != nil {
		t.Left.DisplayPost()
		t.Right.DisplayPost()
		fmt.Printf("%d ", t.Value)
	}
}
func main() {
	// var root *TreeNode
	// root = root.Insert(5)
	root := &TreeNode{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(88)
	fmt.Print("In-Order Traversal :-")
	root.DisplayIn()
	fmt.Println()
	deleteValue := 15
	root = root.Delete(deleteValue)
	fmt.Printf("After deleting %d:\n", deleteValue)
	root.DisplayIn()
	fmt.Println()

	fmt.Print("In-Order Traversal :-")
	root.DisplayIn()
	fmt.Println()
	fmt.Print("Pre-Order Traversal :-")
	root.DisplayPre()
	fmt.Println()
	fmt.Print("Post-Order Traversal :-")
	root.DisplayPost()
	fmt.Println()
	fmt.Println(root.Search(7))
	fmt.Println(root.Search(55))
	root.DisplayBFS()

}
