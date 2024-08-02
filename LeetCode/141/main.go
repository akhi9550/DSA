package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	node4 := &ListNode{Val: -4}
	node3 := &ListNode{Val: 0, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head1 := &ListNode{Val: 3, Next: node2}
	node4.Next = node2 // Creating a cycle here

	fmt.Println(hasCycle(head1)) // Output: true

	node2b := &ListNode{Val: 2}
	head2 := &ListNode{Val: 1, Next: node2b}
	node2b.Next = head2 

	fmt.Println(hasCycle(head2)) // Output: true

	head3 := &ListNode{Val: 1}

	fmt.Println(hasCycle(head3)) // Output: false
}
