package main

import (
	"fmt"
)

type CircularQueue struct {
	items []int
	size  int
	front int
	rear  int
}

func NewCircularQueue(size int) *CircularQueue {
	return &CircularQueue{
		items: make([]int, size),
		size:  size,
		front: -1,
		rear:  -1,
	}
}

func (cq *CircularQueue) enqueue(item int) {
	if cq.isFull() {
		fmt.Println("Queue is full. Cannot enqueue.")
		return
	}

	if cq.isEmpty() {
		cq.front = 0
		cq.rear = 0
	} else {
		cq.rear = (cq.rear + 1) % cq.size
	}

	cq.items[cq.rear] = item
}

func (cq *CircularQueue) dequeue() int {
	if cq.isEmpty() {
		fmt.Println("Queue is empty. Cannot dequeue.")
		return -1 
	}

	item := cq.items[cq.front]

	if cq.front == cq.rear {
		cq.front = -1
		cq.rear = -1
	} else {
		cq.front = (cq.front + 1) % cq.size
	}

	return item
}

func (cq *CircularQueue) isFull() bool {
	return (cq.rear+1)%cq.size == cq.front
}

func (cq *CircularQueue) isEmpty() bool {
	return cq.front == -1 && cq.rear == -1
}

func main() {
	cq := NewCircularQueue(5)

	cq.enqueue(1)
	cq.enqueue(2)
	cq.enqueue(3)
	cq.enqueue(4)
	cq.enqueue(5)

	fmt.Println("Circular Queue after enqueuing elements:")
	printCircularQueue(cq)

	dequeuedItem := cq.dequeue()
	fmt.Printf("Dequeued item: %v\n", dequeuedItem)

	fmt.Println("Circular Queue after dequeuing one element:")
	printCircularQueue(cq)
}

func printCircularQueue(cq *CircularQueue) {
	if cq.isEmpty() {
		fmt.Println("Queue is empty.")
		return
	}

	current := cq.front
	for {
		fmt.Printf("%v ", cq.items[current])

		if current == cq.rear {
			break
		}

		current = (current + 1) % cq.size
	}

	fmt.Println()
}
