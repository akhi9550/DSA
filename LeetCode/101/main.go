package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return helper(root, root)
}

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && helper(left.Right, right.Left) && helper(left.Left, right.Right)
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println("Is the tree symmetric?", isSymmetric(root)) // Output: true

	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println("Is the tree symmetric?", isSymmetric(root2)) // Output: false
}
