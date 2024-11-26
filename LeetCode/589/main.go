package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	arr := []int{root.Val}
	for i := 0; i < len(root.Children); i++ {
		arr = append(arr, preorder(root.Children[i])...)
	}
	return arr
}

func newNode(val int) *Node {
	return &Node{Val: val, Children: []*Node{}}
}

func main() {

	root := newNode(1)
	child1 := newNode(3)
	child2 := newNode(2)
	child3 := newNode(4)

	root.Children = append(root.Children, child1, child2, child3)
	child1.Children = append(child1.Children, newNode(5), newNode(6))
	result := preorder(root)
	fmt.Println("Preorder traversal:", result)
}
