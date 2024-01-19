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

	if heap.Arr[parent] > heap.Arr[index] {
		heap.Arr[parent], heap.Arr[index] = heap.Arr[index], heap.Arr[parent]
	} else {
		return
	}

	heap.HeapifyUp(parent)
}

func (heap *Heap) HeapifyDown(index int){
	for {
		left:=(2*index)+1
		right:=2*index+1
		largest := index

		if left < len(heap.Arr) && heap.Arr[left] > heap.Arr[largest]{
			largest = left
		}

		if right < len(heap.Arr) && heap.Arr[right] > heap.Arr[largest]{
			largest = right
		}
	}
}

func (heap *Heap) remove(val int) int{
	if len(heap.Arr) == 0{
		return 0
	}

	top:=heap.Arr[0]
	heap.Arr[0]=heap.Arr[len(heap.Arr)-1]
	heap.HeapifyDown(0)

	return top 
}

func main() {
	heap := Heap{}
	arr := []int{8,2,4,51,3,23,5,89}

	for _, arrr := range arr {
		heap.Insert(arrr)
	}
	fmt.Println(heap.Arr)

}
