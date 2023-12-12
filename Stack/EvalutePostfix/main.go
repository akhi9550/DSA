package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(i int) {
	s.data = append(s.data, i)
}
func (s *Stack) Pop() int {
	toRemove := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return toRemove
}
func EvalutePostFix(expr string) int {
	stack := Stack{}
	for _, v := range expr {
		if v >= '0' && v <= '9' {
			digit, _ := strconv.Atoi(string(v))
			stack.Push(digit)
		} else {
			operand1 := stack.Pop()
			operand2 := stack.Pop()
			switch v {
			case '+':
				stack.Push(operand1 + operand2)
			case '-':
				stack.Push(operand1 - operand2)
			case '*':
				stack.Push(operand1 * operand2)
			case '/':
				stack.Push(operand1 / operand2)
			}
		}
	}
	return stack.Pop()
}
func main() {
	fmt.Println(EvalutePostFix("22*4+"))   // (2*2) + 4 = 8
	fmt.Println(EvalutePostFix("64+2*5+")) // ((6+4)*2)+5 =25
}
