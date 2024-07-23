package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func findMiddle(head *ListNode) *ListNode {
	fast, slow := head, head
	if fast == nil || fast.Next == nil {
		return slow
	}

	fast = fast.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	middle := findMiddle(head)
	secondHalf := reverseList(middle.Next)
	middle.Next = nil

	firstHalf := head
	for firstHalf != nil && secondHalf != nil {
		temp1 := firstHalf.Next
		temp2 := secondHalf.Next

		firstHalf.Next = secondHalf
		secondHalf.Next = temp1

		firstHalf = temp1
		secondHalf = temp2
	}
}

// Helper function to print the list
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

// Helper function to create a linked list from a slice
func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func main() {
	values := []int{1, 2, 3, 4, 5}
	head := createList(values)

	fmt.Println("Original list:")
	printList(head)

	reorderList(head)

	fmt.Println("Reordered list:")
	printList(head)
}
