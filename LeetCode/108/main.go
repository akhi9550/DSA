package main

import (
    "fmt"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }

    mid := len(nums) / 2
    root := &TreeNode{Val: nums[mid]}

    root.Left = sortedArrayToBST(nums[:mid])
    root.Right = sortedArrayToBST(nums[mid+1:])

    return root
}

func inorderTraversal(root *TreeNode) []int {
    result := []int{}
    inorder(root, &result)
    return result
}

func inorder(root *TreeNode, result *[]int) {
    if root == nil {
        return
    }
    inorder(root.Left, result)
    *result = append(*result, root.Val)
    inorder(root.Right, result)
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7}
    root := sortedArrayToBST(nums)
    result := inorderTraversal(root)
    fmt.Println("In-order Traversal of the BST:", result) // Output: In-order Traversal of the BST: [1 2 3 4 5 6 7]
}
