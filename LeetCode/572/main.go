package main

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
    if s == nil {
        return false
    }

    if s.Val == t.Val {
        if isSame(s, t) {
            return true
        }
    }

    return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(root1 *TreeNode, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {
        return true
    }

    if root1 == nil || root2 == nil {
        return false
    }

    if root1.Val != root2.Val {
        return false
    }

    return isSame(root1.Left, root2.Left) && isSame(root1.Right, root2.Right)
}

func newNode(val int) *TreeNode {
    return &TreeNode{Val: val}
}

func main() {
    s := newNode(3)
    s.Left = newNode(4)
    s.Right = newNode(5)
    s.Left.Left = newNode(1)
    s.Left.Right = newNode(2)
    t := newNode(4)
    t.Left = newNode(1)
    t.Right = newNode(2)

    if isSubtree(s, t) {
        fmt.Println("Tree t is a subtree of tree s")
    } else {
        fmt.Println("Tree t is NOT a subtree of tree s")
    }
}
