package main

import "fmt"

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

// Display performs an in-order traversal and prints the values of the tree.
func (t *TreeNode) DisplayIn() {
	if t != nil {
		t.Left.DisplayIn()
		fmt.Printf("%d ", t.Value)
		t.Right.DisplayIn()
	}
}

// Display performs a preorder traversal and prints the values of the tree.
func (t *TreeNode) DisplayPre() {
	if t != nil {
		fmt.Printf("%d ", t.Value)
		t.Left.DisplayPre()
		t.Right.DisplayPre()
	}
}

// Display performs a post-order traversal and prints the values of the tree.
func (t *TreeNode) DisplayPost() {
	if t != nil {
		t.Left.DisplayPost()
		t.Right.DisplayPost()
		fmt.Printf("%d ", t.Value)
	}
}
func main() {
	root := &TreeNode{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(88)
	root.DisplayIn()
	fmt.Println()
	root.DisplayPre()
	fmt.Println()
	root.DisplayPost()
	fmt.Println()
	fmt.Println(root.Search(7))
	fmt.Println(root.Search(55))

}
