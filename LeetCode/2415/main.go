package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	level := 0
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		if level%2 == 1 {
			l, r := 0, len(nodes)-1
			for l < r {
				nodes[l].Val, nodes[r].Val = nodes[r].Val, nodes[l].Val
				l++
				r--
			}
		}

		length := len(nodes)
		for i := 0; i < length; i++ {
			if nodes[i].Left == nil {
				break
			}
			nodes = append(nodes, nodes[i].Left, nodes[i].Right)
		}
		nodes = nodes[length:]
		level++
	}
	return root
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	root = reverseOddLevels(root)

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := []*TreeNode{}
		for _, node := range queue {
			fmt.Print(node.Val, " ")
			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
		fmt.Println()
		queue = level
	}
}
