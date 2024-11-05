package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 != nil && root2 != nil {
		root1.Val += root2.Val
	} else if root1 != nil && root2 == nil {
		return root1
	} else if root1 == nil && root2 != nil {
		return root2
	}

	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)

	return root1
}

// Helper function to print the tree in order (for testing purposes).
func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	printTree(root.Left)
	fmt.Print(root.Val, " ")
	printTree(root.Right)
}

func main() {
	// Creating first tree.
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 3}
	root1.Right = &TreeNode{Val: 2}
	root1.Left.Left = &TreeNode{Val: 5}

	// Creating second tree.
	root2 := &TreeNode{Val: 2}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 3}
	root2.Left.Right = &TreeNode{Val: 4}
	root2.Right.Right = &TreeNode{Val: 7}

	mergedTree := mergeTrees(root1, root2)

	fmt.Print("Merged Tree (In-order): ")
	printTree(mergedTree)
	fmt.Println()
}
