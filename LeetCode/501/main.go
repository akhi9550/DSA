package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	m := make(map[int]int)
	inorder(root, m)
	var mode int
	for _, count := range m {
		if mode <= count {
			mode = count
		}
	}

	result := make([]int, 0)
	for num, count := range m {
		if count == mode {
			result = append(result, num)
		}
	}

	return result
}

func inorder(root *TreeNode, m map[int]int) {
	if root != nil {
		inorder(root.Left, m)
		m[root.Val]++
		inorder(root.Right, m)
	}
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else if val > root.Val {
		root.Right = insert(root.Right, val)
	}
	return root
}

func main() {
	root := &TreeNode{Val: 5}
	insert(root, 3)
	insert(root, 8)
	insert(root, 3) // duplicate value
	insert(root, 2)
	insert(root, 5) // duplicate value
	insert(root, 7)
	modes := findMode(root)
	fmt.Println("Mode(s):", modes)
}
