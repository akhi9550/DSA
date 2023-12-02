package main

import "fmt"

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
	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

func (list *LinkedList) Print() {
	temp := list.head
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
	fmt.Println()
}

func (list *LinkedList) Reverse() {
	var prev, next *Node
	temp := list.head
	for temp != nil {
		next = temp.next
		temp.next = prev
		prev = temp
		temp = next
	}
	list.head = prev
}

func main() {
	mylist := LinkedList{}
	mylist.Append(2)
	mylist.Append(3)
	mylist.Append(9)
	fmt.Println("List before reversing:")
	mylist.Print()
	mylist.Reverse()
	fmt.Println("List after reversing:")
	mylist.Print()
}
