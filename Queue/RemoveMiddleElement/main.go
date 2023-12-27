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
	firstelement := q.arr[0]
	q.arr = q.arr[1:]
	return firstelement
}
func (q *queue) removeMiddle(midindex int, currentindex int) {

	if currentindex == midindex {
		mid := q.dequeue()
		fmt.Println("mid element is ", mid)
		return
	}
	temp := q.dequeue()
	q.removeMiddle(midindex, currentindex+1)
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
	currentindex := 0
	midindex := q.size / 2
	fmt.Println(q.arr)
	q.removeMiddle(midindex, currentindex)
	fmt.Println(q.arr)
}
