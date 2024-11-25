package main

import (
	"fmt"
)

type KthLargest struct {
	mh *MinHeap
	k  int
}

func Constructor(k int, nums []int) KthLargest {
	mh := NewMinHeap()
	for i := 0; i < len(nums); i++ {
		mh.Add(nums[i])
	}

	for mh.Len() > k {
		mh.Pop()
	}

	return KthLargest{mh, k}
}

func (this *KthLargest) Add(val int) int {
	this.mh.Add(val)
	for this.mh.Len() > this.k {
		this.mh.Pop()
	}

	return this.mh.Top()
}

type MinHeap struct {
	items []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (mh *MinHeap) Add(x int) {
	l := 0
	r := len(mh.items) - 1

	for l <= r {
		m := l + (r-l)/2

		if x == mh.items[m] {
			l = m
			break
		} else if x > mh.items[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	if l == len(mh.items) {
		mh.items = append(mh.items, x)
		return
	}

	mh.items = append(mh.items[:l+1], mh.items[l:]...)
	mh.items[l] = x
}

func (mh *MinHeap) Top() int {
	return mh.items[0]
}

func (mh *MinHeap) Pop() {
	mh.items = mh.items[1:]
}

func (mh *MinHeap) Len() int {
	return len(mh.items)
}

func main() {
	k := 3
	nums := []int{4, 5, 8, 2}
	obj := Constructor(k, nums)
	fmt.Println(obj.Add(3))  // Output: 4
	fmt.Println(obj.Add(5))  // Output: 5
	fmt.Println(obj.Add(10)) // Output: 5
	fmt.Println(obj.Add(9))  // Output: 8
	fmt.Println(obj.Add(4))  // Output: 8
}
