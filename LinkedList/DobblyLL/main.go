package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}
type DoubleLL struct {
	head *Node
	tail *Node
}

func (dll *DoubleLL) AddNode(data int) {
	newNode := &Node{data: data, next: nil, prev: nil}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		//began
		// newNode.next = dll.head
		// dll.head.prev = newNode
		// dll.head = newNode
		/////end
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
}
func (dll *DoubleLL) Print() {
	temp := dll.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func (dll *DoubleLL) DeleteAtBeginnig() {
	if dll.head == nil {
		fmt.Println("List Empty")
		return
	}
	if dll.head.next != nil {
		dll.head = dll.head.next
		dll.head.prev = nil
	} else {
		dll.head = nil
		dll.tail = nil
	}
}
func (dll *DoubleLL) DeleteEnd() {
	if dll.head == nil {
		fmt.Println("Empty")
		return
	}
	if dll.head == dll.tail {
		dll.head, dll.tail = nil, nil
		return
	}
	dll.tail = dll.tail.prev
	dll.tail.next = nil
}
func (dll *DoubleLL) Reverse() {
	temp := dll.tail
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.prev
	}
}
func (dll *DoubleLL) Insert(data int) {
	newNode:=&Node{data: data,prev: nil,next: nil}
	if dll.head == nil{
		dll.head = nil
		dll.tail = nil
	}
	//end
	newNode.prev = dll.tail
	dll.tail.next = newNode
	dll.tail = newNode
	//begin
	// dll.head.prev =newNode
	// newNode.next=dll.head
	// dll.head = newNode
	


}
func main() {
	dll := DoubleLL{}
	dll.AddNode(9)
	dll.AddNode(5)
	dll.AddNode(7)
	dll.AddNode(8)
	fmt.Println("Doubly list:")
	dll.Print()
	fmt.Println("insert begin")
	dll.Insert(78)
	dll.Print()
	// fmt.Println("After delete Begin")
	// dll.DeleteAtBeginnig()
	// dll.Print()
	// fmt.Println("After delete End")
	// dll.DeleteEnd()
	// dll.Print()

}
