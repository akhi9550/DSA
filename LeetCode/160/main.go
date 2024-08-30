package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodes := make(map[*ListNode]bool)

	for headA != nil {
		nodes[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if ok := nodes[headB]; ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}

func createLinkedList(values []int) *ListNode {
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

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	listA := createLinkedList([]int{4, 1})
	listB := createLinkedList([]int{5, 6, 1})

	intersection := createLinkedList([]int{8, 4, 5})

	listA.Next.Next = intersection
	listB.Next.Next.Next = intersection

	fmt.Println("List A:")
	printLinkedList(listA)

	fmt.Println("List B:")
	printLinkedList(listB)

	intersectionNode := getIntersectionNode(listA, listB)
	if intersectionNode != nil {
		fmt.Printf("Intersection Node Value: %d\n", intersectionNode.Val)
	} else {
		fmt.Println("No intersection found.")
	}
}
