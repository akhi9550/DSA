package main

import "fmt"

type Heap struct {
	array []int
}

func parent(i int) int {
	return (i - 1) / 2
}
func leftChild(i int) int {
	return (2 * i) + 1
}
func rightChild(i int) int {
	return (2 * i) + 2
}

func (h *Heap) Insert(i int) {
	h.array = append(h.array, i)
}
func (h *Heap) Build() {
	for i := parent(len(h.array) - 1); i >= 0; i-- {
		h.shiftDown(i)
	}
}
func (h *Heap) swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
func (h *Heap) shiftDown(currentIdx int) {
	endIdx := len(h.array) - 1
	var idxToSwap int
	for {
		leftIdx := leftChild(currentIdx)
		rightIdx := rightChild(currentIdx)
		if leftIdx < endIdx {
			if rightIdx <= endIdx && h.array[leftIdx] < h.array[rightIdx] {
				idxToSwap = leftIdx
			} else {
				idxToSwap = rightIdx
			}
			if h.array[idxToSwap] < h.array[currentIdx] {
				h.swap(h.array, idxToSwap, currentIdx)
				currentIdx = idxToSwap
			} else {
				break
			}
		} else {
			break
		}
	}
}

func (h *Heap) Remove() int {
	if len(h.array) == 0 {
		return -1
	}
	minVal := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.shiftDown(0)
	return minVal
}
func (h *Heap) Display() {
	for _, v := range h.array {
		fmt.Println(v, " ")
	}
	fmt.Println()
}
func main() {
	arr := []int{2, 3, 4, 6, 9, 8, 7}
	heap := &Heap{}
	for _, v := range arr {
		heap.Insert(v)
	}
	fmt.Println("Array is :-", arr)
	heap.Build()
	fmt.Println("Heap is :-")
	heap.Display()
	for i := 0; i <= 3; i++ {
		heap.Remove()
	}
	fmt.Println("After Delete :-")
	heap.Display()

}
