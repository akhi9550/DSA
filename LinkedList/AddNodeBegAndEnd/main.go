package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func (list *LinkedList) AddNodeAtBeginnning(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}
	newNode.next = list.head
	list.head = newNode
}
func (list *LinkedList) Print() {
	temp := list.head
	for temp != nil {
		fmt.Println(temp.data, "")
		temp = temp.next
	}
	fmt.Println()
}
func (list *LinkedList) AddNodeAtEnd(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}
	newNode.next = nil
	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}
func main() {
	mylist := LinkedList{}
	mylist.AddNodeAtBeginnning(1)
	mylist.AddNodeAtBeginnning(3)
	fmt.Println("Add beginnig")
	mylist.Print()
	mylist.AddNodeAtEnd(9)
	mylist.AddNodeAtEnd(8)
	fmt.Println("Add End")
	mylist.Print()
}
