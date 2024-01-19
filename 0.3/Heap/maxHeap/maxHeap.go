package main

import "fmt"

type Heap struct {
	Arr []int
}

func (heap *Heap) Insert(val int) {
	heap.Arr = append(heap.Arr, val)
	heap.HeapifyUp(len(heap.Arr) - 1)
}

func (heap *Heap) HeapifyUp(index int) {
	if index == 0 {
		return
	}
	parent := (index - 1) / 2

	if heap.Arr[parent] < heap.Arr[index] {
		heap.Arr[parent], heap.Arr[index] = heap.Arr[index], heap.Arr[parent]
	} else {
		return
	}

	heap.HeapifyUp(parent)
}



func main() {
	heap := Heap{}
	arr := []int{8,2,4,51,3,23,5,89}

	for _, arrr := range arr {
		heap.Insert(arrr)
	}
	fmt.Println(heap.Arr)

}
