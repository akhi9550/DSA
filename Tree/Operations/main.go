package main

import "fmt"

type Node struct {
	Left  *Node
	Value int
	Right *Node 
}
type BinaryTree struct {
	root *Node
}

func main() {
	node1 := &Node{nil, 7, nil}
	node2 := &Node{nil, 3, nil}
	node3 := &Node{nil, 4, nil}
	node4 := &Node{nil, 8, nil}
	node5 := &Node{nil, 6, nil}
	node6 := &Node{nil, 1, nil}
	node7 := &Node{nil, 9, nil}

	b := &BinaryTree{}

	b.root = node1
	b.root.Left = node2
	b.root.Right = node3

	b.root.Left.Left = node4
	b.root.Left.Right = node5

	b.root.Right.Left = node6
	b.root.Right.Right = node7

	fmt.Print("Pre-Order Traversal :-")
	PreOrder(b.root)
	fmt.Println()

	fmt.Print("In-Order Traversal :-")
	InOrder(b.root)
	fmt.Println()

	fmt.Print("Post-Order Traversal :-")
	PostOrder(b.root)
	fmt.Println()

}

//Pre-Order Traversal
func PreOrder(root *Node) {
	if root == nil {
		return
	}
	PreOrder(root.Left)
	PreOrder(root.Right)
	fmt.Print(root.Value, " ")
}

//In-Order Traversal
func InOrder(root *Node) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Print(root.Value, " ")
	InOrder(root.Right)
}

//Post-Order Traversal
func PostOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Value, " ")
	PostOrder(root.Left)
	PostOrder(root.Right)
}
