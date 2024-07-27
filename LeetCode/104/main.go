package main

import (
    "fmt"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    left := maxDepth(root.Left)
    right := maxDepth(root.Right)
    
    if left > right {
        return left + 1
    }
    return right + 1
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
        },
    }

    depth := maxDepth(root)

    fmt.Println("Maximum Depth:", depth) // Output: Maximum Depth: 3
}
