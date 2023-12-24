package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type Stack struct {
	head *Node
	size int
}

func (s *Stack) Push(i int) {
	newNode := &Node{data: i, next: nil}
	if s.head == nil {
		s.head = newNode
		s.size++
	} else {
		newNode.next = s.head
		s.head = newNode
		s.size++
	}
}
func (s *Stack) Pop() {
	value := s.head.data
	s.head = s.head.next
	fmt.Println("Popped Element is ", value)
}
func (s *Stack) Display() {
	temp := s.head
	if temp == nil {
		fmt.Println("Stack Empty")
		return
	}
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}
func (s *Stack) Size() int {
	return s.size
}
func (s *Stack) IsEmpty() {
	if s.head == nil {
		fmt.Println("stack Empty")
	} else {
		fmt.Println("stack is not Empty")
	}
}
func (s *Stack) Top() {
	if s.head == nil {
		fmt.Println("stack empty")
	} else {
		fmt.Println("Stack Top Element is :", s.head.data)
	}
}
func main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(23)
	s.Display()
	fmt.Println()
	fmt.Println("Size of Stack:", s.Size())
	s.IsEmpty()
	s.Top()
	s.Pop()
	fmt.Println("Final Stack Elements:")
	s.Display()
}
