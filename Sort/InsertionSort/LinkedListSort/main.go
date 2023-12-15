package main

import (
	"fmt"
)

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func main() {
	head := &LinkedList{Val: 4, Next: &LinkedList{Val: 2, Next: &LinkedList{Val: 7, Next: &LinkedList{Val: 5}}}}
	fmt.Println("The unsorted linked list")
	Display(head)
	sortedList := LinkedListSorting(head)
	fmt.Println("The sorted linked LIst:")
	Display(sortedList)
}
func LinkedListSorting(head *LinkedList) *LinkedList {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &LinkedList{Val: 0, Next: nil}
	current := head
	for current != nil {
		prev, temp := dummy, dummy.Next
		next := current.Next
		for temp != nil && temp.Val < current.Val {
			prev = temp
			temp = temp.Next
		}
		prev.Next = current
		current.Next = temp
		current = next
	}
	return dummy.Next

}
func Display(head *LinkedList) {
	current := head
	for current != nil {
		fmt.Printf("%d", current.Val)
		current = current.Next
	}
	fmt.Println()
}