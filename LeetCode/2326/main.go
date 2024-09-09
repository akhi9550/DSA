package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
    mat := make([][]int, m)
    for i := 0; i < m; i++ {
        mat[i] = make([]int, n)
        for j := 0; j < n; j++ {
            mat[i][j] = -1 
        }
    }
    y, x := 0, 0
    dirs := map[int][]int{
        1: {0, 1},  
        2: {1, 0},  
        3: {0, -1}, 
        4: {-1, 0}, 
    }
    d, shrink, lim := 1, 0, 0

    for head != nil {
        mat[y][x] = head.Val
        head = head.Next
        i, j := y+dirs[d][0], x+dirs[d][1]
        if i >= shrink && i < m-lim && j >= lim && j < n-lim {
            y, x = i, j
        } else {
            d++
            if d == 4 {
                shrink++
            } else if d > 4 {
                d = 1
                lim++
            }
            y, x = y+dirs[d][0], x+dirs[d][1]
        }
    }
    return mat
}

func createLinkedList(nums []int) *ListNode {
    if len(nums) == 0 {
        return nil
    }
    head := &ListNode{Val: nums[0]}
    current := head
    for _, num := range nums[1:] {
        current.Next = &ListNode{Val: num}
        current = current.Next
    }
    return head
}

func main() {
    m, n := 3, 4

    listValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
    head := createLinkedList(listValues)

    matrix := spiralMatrix(m, n, head)
    for _, row := range matrix {
        fmt.Println(row)
    }
}
