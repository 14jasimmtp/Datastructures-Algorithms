package main

import "fmt"

type Heap struct {
	Arr []int
}

func getParentIndex(index int) int{
	return (index-1)/2
}

func getLeftChildIndex(index int) int{
	return 2*index+1
}

func getRightChildIndex(index int) int{
	return 2*index+2
}

func (h *Heap)swap(index1 ,index2 int){
	h.Arr[index1],h.Arr[index2] = h.Arr[index2],h.Arr[index1]
}


func (heap *Heap) Insert(val int) {
	heap.Arr = append(heap.Arr, val)
	heap.HeapifyUp(len(heap.Arr) - 1)
}


func (heap *Heap) HeapifyDown(index int){
	leftchild:=getLeftChildIndex(index)
	rightchild:=getRightChildIndex(index)
	maxIndex:= index
	if leftchild < len(heap.Arr) && heap.Arr[leftchild] > heap.Arr[index]{
		maxIndex = leftchild
	}

	if rightchild < len(heap.Arr) && heap.Arr[rightchild] > heap.Arr[index]{
		maxIndex = rightchild
	}

	if maxIndex != index{
		
	}
}
func (heap *Heap) HeapifyUp(index int) {
	if index == 0 {
		return
	}
	parent := (index - 1) / 2

	if parent >= 0 && heap.Arr[parent] < heap.Arr[index] {
		heap.Arr[parent], heap.Arr[index] = heap.Arr[index], heap.Arr[parent]
	} 
	heap.HeapifyUp(parent)
}



func main() {
	heap := Heap{}
	arr := []int{8, 2, 4, 51, 3, 23, 5, 89}

	for _, arrr := range arr {
		heap.Insert(arrr)
	}
	fmt.Println(heap.Arr)

}
