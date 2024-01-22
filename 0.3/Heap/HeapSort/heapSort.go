package main

import "fmt"

type Heap struct {
	arr []int
}

func (h *Heap) heapifyUP(index int) {
	parent := (index - 1) / 2

	if parent >= 0 && h.arr[parent] < h.arr[index] {
		h.arr[index], h.arr[parent] = h.arr[parent], h.arr[index]
		h.heapifyUP(parent)
	}

}

func (h *Heap) heapifyDown(index int) {
	leftchild := 2*index + 1
	rightchild := 2*index + 2
	maxIndex := index
	if leftchild < len(h.arr) && h.arr[leftchild] > h.arr[maxIndex] {
		maxIndex = leftchild
	}

	if rightchild < len(h.arr) && h.arr[rightchild] > h.arr[maxIndex] {
		maxIndex = rightchild
	}

	if maxIndex != index {
		h.arr[index], h.arr[maxIndex] = h.arr[maxIndex], h.arr[index]
		h.heapifyDown(maxIndex)
	}

}

func (h *Heap) buildHeapFromArray(array []int) {
	parent := (len(array) - 2) / 2

	for i := parent; i <= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h *Heap) insert(val int) {
	h.arr = append(h.arr, val)
	h.heapifyUP(len(h.arr)-1)
}

func (h *Heap) remove() int {
	if len(h.arr) < 1 {
		return -1
	}

	if len(h.arr) == 1 {
		rem := h.arr[0]
		h.arr = h.arr[1:]
		return rem
	}
	rem := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.heapifyDown(0)
	return rem

}

func (h *Heap) heapSort() []int {
	sortedArray := []int{}
	for len(h.arr) > 0 {
		sortedArray = append(sortedArray, h.remove())
	}
	return sortedArray
}

func main() {
	h := Heap{}
	h.insert(9)
	h.insert(4)
	h.insert(6)
	h.remove()
	h.insert(7)
	fmt.Println(h.heapSort())

}
