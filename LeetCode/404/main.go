package main

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
    var resp = TreeNode{
        Val: 0,
    }
    add(root, &resp)
    return resp.Val
}

func add(node, root *TreeNode) {
    if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
        root.Val += node.Left.Val
    }
    if node.Right != nil {
        add(node.Right, root)
    }
    if node.Left != nil {
        add(node.Left, root)
    }
    return
}

func newTreeNode(val int) *TreeNode {
    return &TreeNode{Val: val}
}

func main() {
    root := newTreeNode(3)
    root.Left = newTreeNode(9)
    root.Right = newTreeNode(20)
    root.Right.Left = newTreeNode(15)
    root.Right.Right = newTreeNode(7)

    result := sumOfLeftLeaves(root)

    fmt.Println("Sum of left leaves:", result)
}
