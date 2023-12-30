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
func main() {
	root := &TreeNode{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(88)
	fmt.Println(root.Search(7))
	fmt.Println(root.Search(55))
}
