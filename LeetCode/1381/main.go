package main

import (
	"fmt"
)

type CustomStack struct {
	stack   []int
	maxSize int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stack:   make([]int, 0),
		maxSize: maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.stack) < this.maxSize {
		this.stack = append(this.stack, x)
	}
}

func (this *CustomStack) Pop() int {
	val := -1

	if len(this.stack) > 0 {
		val = this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
	}

	return val
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < min(len(this.stack), k); i++ {
		this.stack[i] += val
	}
}

func min(nums ...int) int {
	minNum := nums[0]
	for _, num := range nums[1:] {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}

func main() {
	stack := Constructor(3) 

	stack.Push(1)          
	stack.Push(2)          
	fmt.Println(stack.Pop()) 

	stack.Push(2)        
	stack.Push(3)         
	stack.Push(4)           
	stack.Increment(5, 100) 
	stack.Increment(2, 100) 

	fmt.Println(stack.Pop()) 
	fmt.Println(stack.Pop()) 
	fmt.Println(stack.Pop()) 
	fmt.Println(stack.Pop()) 
}
