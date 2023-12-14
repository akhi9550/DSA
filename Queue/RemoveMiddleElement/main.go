package main

import "fmt"

type queue struct {
	arr  []int
	size int
}

func (q *queue) enqueue(data int) {

	q.arr = append(q.arr, data)
	q.size++
}
func (q *queue) dequeue() int {
	frstelement := q.arr[0]
	q.arr = q.arr[1:]
	return frstelement
}
func (q *queue) removeMiddle(mididx int, currentidx int) {

	if currentidx == mididx {
		mid := q.dequeue()
		fmt.Println("mid element is ", mid)
		return
	}
	temp := q.dequeue()
	q.removeMiddle(mididx, currentidx+1)
	q.enqueue(temp)

}

func main() {
	q := queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(8)
	q.enqueue(3)
	q.enqueue(4)
	q.enqueue(5)	
	currentidx := 0
	mididx := q.size / 2
	fmt.Println(q.arr)
	q.removeMiddle(mididx, currentidx)
	fmt.Println(q.arr)
}