package main

import "fmt"

type stack struct {
	arr  []int
	size int
}

func (s *stack) push(data int) {
	s.arr = append(s.arr, data)
	s.size++
}
func (s *stack) pop() int {
	if s.size == 0 {
		return 0
	}
	lastElement := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return lastElement
}
func (s *stack) Display() {
	fmt.Println("Elements are : ", s.arr)
}
func (s *stack) removeMiddle(middleidx int, currentIdx int) int {
	if middleidx == currentIdx {
		return s.pop()
	}
	temp := s.pop()
	mid := s.removeMiddle(middleidx, currentIdx+1)
	s.push(temp)
	return mid
}
func main() {
	s := &stack{}
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(7)
	s.push(6)
	s.push(4)
	s.push(5)
	s.Display()
	middleidx := s.size / 2
	currentIdx := 0
	fmt.Println(s.removeMiddle(middleidx, currentIdx))
}
