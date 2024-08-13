package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var previous *ListNode = nil
	temp := head
	for temp != nil {
		next := temp.Next
		temp.Next = previous
		previous = temp
		temp = next
	}
	return previous
}

func printList(head *ListNode) {
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
	}
	fmt.Println()
}

func main() {
	// Create a linked list: 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Println("Original List:")
	printList(head)

	reversedHead := reverseList(head)

	fmt.Println("Reversed List:")
	printList(reversedHead)
}
