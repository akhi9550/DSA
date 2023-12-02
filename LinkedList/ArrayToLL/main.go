package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func (list *LinkedList) ArrayConvert(arr []int) {
	for _, i := range arr {
		newNode := &Node{data: i, next: nil}
		if list.head == nil {
			list.head = newNode
		} else { 
			temp := list.head
			for temp.next != nil {
				temp = temp.next
			}
			temp.next = newNode
		}
	}
}
func (list *LinkedList) Print() {
	temp := list.head
	for temp != nil {
		fmt.Printf("%d ",temp.data)
		temp = temp.next
	}

}

func main() {
	list := LinkedList{}
	arr := []int{21, 22, 23, 24, 25, 27}
	fmt.Println("Array:", arr)
	list.ArrayConvert(arr)
	fmt.Print("LinkedList:")
	list.Print()
}

// func ToConvertArrayToLL(arr []int) *Node {

// }
// func PrintLinkedList(head *Node) {
// 	for head != nil {
// 		fmt.Printf("%d ", head.data)
// 		head = head.next
// 	}
// 	fmt.Println("")
// }
// 	var head, tail *Node
// 	for _, val := range arr {
// 		newNode := &Node{data: val,next: nil}
// 		if head == nil {
// 			head = newNode
// 			tail = newNode
// 		} else {
// 			tail.next = newNode
// 			tail = newNode
// 		}
// 	}
// 	return head
