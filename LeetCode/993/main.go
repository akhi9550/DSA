package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func parentsAndDepth(x int, root *TreeNode) (*TreeNode, int) {
	switch {
	case root == nil:
		return nil, -1
	case (root.Left != nil && root.Left.Val == x) || (root.Right != nil && root.Right.Val == x):
		return root, 1
	}

	if p, l := parentsAndDepth(x, root.Left); l > 0 {
		return p, l + 1
	}

	if p, r := parentsAndDepth(x, root.Right); r > 0 {
		return p, r + 1
	}

	return nil, -1
}

func isCousins(root *TreeNode, x int, y int) bool {
	xp, xd := parentsAndDepth(x, root)
	yp, yd := parentsAndDepth(y, root)

	return xd == yd && xp != yp
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	fmt.Println(isCousins(root, 4, 5)) // Output: true
	fmt.Println(isCousins(root, 4, 3)) // Output: false
}
