package main

import "fmt"

type Queue struct {
	head *Node
	end  *Node
}
type Node struct {
	data int
	next *Node
}

func (q *Queue) enqueue(element int) {
	newNode := &Node{data: element, next: nil}
	if q.head == nil {
		q.head = newNode
	} else {
		q.end.next = newNode
	}
	q.end = newNode
}
func (q *Queue) dequeue() {
	if q.head == nil {
		return
	}
	q.head = q.head.next
}
func (q *Queue) Display() {
	if q.head == nil {
		fmt.Println("Nothing to print")
		return
	}
	temp := q.head
	for temp != nil {
		fmt.Printf(" %d",temp.data)
		temp = temp.next
	}
	fmt.Println(" ")
}
func main() {
	que := Queue{}
	que.enqueue(1)
	que.enqueue(2)
	que.enqueue(3)
	que.enqueue(4)
	que.Display()
	que.dequeue()
	que.Display()
}
