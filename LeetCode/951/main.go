package main

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {
        return true
    }
    if root1 == nil || root2 == nil {
        return false
    }
    if root1.Val != root2.Val {
        return false
    }
    return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
           (flipEquiv(root1.Right, root2.Left) && flipEquiv(root1.Left, root2.Right))
}

func main() {
    root1 := &TreeNode{Val: 1}
    root1.Left = &TreeNode{Val: 2}
    root1.Right = &TreeNode{Val: 3}
    root1.Left.Left = &TreeNode{Val: 4}
    root1.Left.Right = &TreeNode{Val: 5}
    root1.Right.Left = &TreeNode{Val: 6}
    root1.Left.Right.Left = &TreeNode{Val: 7}
    root1.Left.Right.Right = &TreeNode{Val: 8}

    root2 := &TreeNode{Val: 1}
    root2.Left = &TreeNode{Val: 3}
    root2.Right = &TreeNode{Val: 2}
    root2.Left.Right = &TreeNode{Val: 6}
    root2.Right.Left = &TreeNode{Val: 4}
    root2.Right.Right = &TreeNode{Val: 5}
    root2.Right.Right.Left = &TreeNode{Val: 8}
    root2.Right.Right.Right = &TreeNode{Val: 7}
    result := flipEquiv(root1, root2)
    if result {
        fmt.Println("The trees are flip equivalent.")
    } else {
        fmt.Println("The trees are not flip equivalent.")
    }
}
