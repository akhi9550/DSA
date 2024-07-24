package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	s := []int{}
	head := &ListNode{}

	for _, v := range lists {
		for v != nil {
			s = append(s, v.Val)
			v = v.Next
		}
	}

	cur := head

	slices.Sort(s)

	for _, v := range s {
		temp := &ListNode{Val: v}
		cur.Next = temp
		cur = cur.Next
	}

	return head.Next
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Println()
}

func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	current := head
	for _, val := range vals[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func main() {
	list1 := createList([]int{1, 4, 5})
	list2 := createList([]int{1, 3, 4})
	list3 := createList([]int{2, 6})
	lists := []*ListNode{list1, list2, list3}

	mergedList := mergeKLists(lists)

	fmt.Println("Merged List:")
	printList(mergedList)
}
