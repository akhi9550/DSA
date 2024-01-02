package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) insert(data int) {
	h.array = append(h.array, data)
	h.BuildHeap()
}
func (h *MaxHeap) BuildHeap() {
	n := len(h.array)
	for i := n/2 - 1; i >= 0; i-- {
		h.shiftDown(i)
	}
}
func (h *MaxHeap) shiftDown(currentIdx int) {
	endIdx := len(h.array) - 1
	for {
		leftIdx := leftChild(currentIdx)
		if leftIdx > endIdx {
			break
		}
		var idxToSwap int
		rightIdx := rightChild(currentIdx)
		if rightIdx <= endIdx && h.array[rightIdx] > h.array[leftIdx] {
			idxToSwap = rightIdx
		} else {
			idxToSwap = leftIdx
		}
		if h.array[idxToSwap] > h.array[currentIdx] {
			swap(h.array, idxToSwap, currentIdx)
			currentIdx = idxToSwap
		} else {
			break
		}
	}
}
func (h *MaxHeap) Delete() int {
	if len(h.array) == 0 {
		return -1
	}
	maxValue := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.shiftDown(0)
	return maxValue
}
func (h *MaxHeap) heapSort() []int {
	n := len(h.array)
	for i := n - 1; i > 0; i-- {
		h.array[0], h.array[i] = h.array[i], h.array[0]
		h.shiftDown(0)
	}
	return h.array
}
func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

func swap(arr []int, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func display(arr []int) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println(" ")
}
func main() {
	var arr = []int{11, 2, 7, 4, 12, 5, 9, 10}
	h := &MaxHeap{}

	for _, v := range arr {
		h.insert(v)
	}

	fmt.Println("Original array:")
	display(h.array)

	h.heapSort()

	fmt.Println("Sorted array:")
	display(h.array)
}
