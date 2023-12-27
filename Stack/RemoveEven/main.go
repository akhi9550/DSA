package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Push(i int) {
	s.data = append(s.data, i)
}
func (s *Stack) Pop() int {
	n := len(s.data) - 1
	toRemove := s.data[n]
	s.data = s.data[:n]
	return toRemove
}
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func RemoveEvenNumbers(s *Stack) {
	if s.IsEmpty() {
		return
	}
	toRemove := s.Pop()
	RemoveEvenNumbers(s)
	if toRemove%2 != 0 {
		s.Push(toRemove)
	}
}
func main() {
	Mystack := &Stack{}
	Mystack.Push(2)
	Mystack.Push(3)
	Mystack.Push(4)
	Mystack.Push(5)
	Mystack.Push(7)
	Mystack.Push(69)
	Mystack.Push(17)
	Mystack.Push(70)
	Mystack.Push(72)
	fmt.Println("The Array:", Mystack.data)
	RemoveEvenNumbers(Mystack)
	if Mystack.IsEmpty() {
		fmt.Println("The array is empty")
	}
	fmt.Println("Array after the removal of even numbers:", Mystack.data)
}
