package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	n, curr := 0, head
	for curr != nil {
		n += 1
		curr = curr.Next
	}

	m, r := n/k, n%k

	res := make([]*ListNode, k)
	curr = head
	for j := 0; j < k; j++ {
		res[j] = curr
		partSize := m
		if r > 0 {
			partSize++
			r--
		}
		for i := 1; i < partSize && curr != nil; i++ {
			curr = curr.Next
		}
		if curr != nil {
			tmp := curr.Next
			curr.Next = nil
			curr = tmp
		}
	}

	return res
}

func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	curr := head
	for _, val := range arr[1:] {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	head := createList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	k := 3
	parts := splitListToParts(head, k)
	for i, part := range parts {
		fmt.Printf("Part %d: ", i+1)
		printList(part)
	}
}
