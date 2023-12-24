package main

import (
	"fmt"
	"strconv"
)

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}
type bucket struct {
	head *backetNode
}
type backetNode struct {
	data int
	next *backetNode
}

func (h *HashTable) Insert(i int) {
	index := Hash(i)
	h.array[index].insert(i)
}

func (h *HashTable) Search(i int) bool {
	index := Hash(i)
	return h.array[index].search(i)
}

func (h *HashTable) Delete(i int) {
	index := Hash(i)
	h.array[index].delete(i)
}
func (b *bucket) insert(i int) {
	if !b.search(i) {
		newNode := &backetNode{data: i, next: nil}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Already Exist")
	}
}
func (b *bucket) search(i int) bool {
	temp := b.head
	for temp != nil {
		if temp.data == i {
			return true
		}
		temp = temp.next
	}
	return false
}
func (b *bucket) delete(i int) {
	if b.head.data == i {
		b.head = b.head.next
		return

	}
	prevNode := b.head
	for prevNode.next != nil {
		prevNode.next = prevNode.next.next
	}
	prevNode = prevNode.next
}
func Hash(i int) int {
	sum := 0
	for _, c := range strconv.Itoa(i) {
		sum += int(c)
	}
	return sum % ArraySize
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func (h *HashTable) Display() {
	for i, v := range h.array {
		fmt.Printf("Buckent %d: ", i)
		temp := v.head
		for temp != nil {
			fmt.Printf("%d ", temp.data)
			temp = temp.next
		}
		fmt.Println()
	}
}

func main() {
	table := Init()
	digits := []int{1, 2, 43, 23, 53, 564, 987}
	for _, v := range digits {
		table.Insert(v)
	}
	table.Delete(53)
	fmt.Println("Search 53", table.Search(53))
	fmt.Println("Search 23", table.Search(23))
	table.Display()
}
