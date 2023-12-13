package main

import "fmt"

func QueneToStack(queue []int) map[int]int {
	stack := make(map[int]int)
	for i, v := range queue {
		stack[i] = v
	}
	return stack
}
func main() {
	queue := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("the quene:", queue)
	stack := QueneToStack(queue)
	fmt.Println("the stack", stack)
}
