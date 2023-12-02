package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func (list *LinkedList) Append(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}
	/////--Add begin--/////
	// newNode.next = list.head
	// list.head = newNode
	////////////////////////

	/////------Add End------///////

	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}
func (list *LinkedList) InsertAtPosition(data, position int) {
	newNode := &Node{data: data, next: nil}
	if position == 1 {
		newNode.next = list.head
		list.head = newNode
		return
	}
	temp := list.head
	for i := 1; i < position-1 && temp != nil; i++ {
		temp = temp.next
	}
	if temp == nil {
		fmt.Println("invalid")
		return
	}
	newNode.next = temp.next
	temp.next = newNode
}
func (list *LinkedList) Print() {
	temp := list.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func main() {
	mylist := LinkedList{}
	mylist.Append(1)
	mylist.Append(4)
	mylist.Append(6)
	fmt.Println("Linked List :")
	mylist.Print()
	mylist.InsertAtPosition(2, 2)
	fmt.Println("Position Inserted :")
	mylist.Print()

}
