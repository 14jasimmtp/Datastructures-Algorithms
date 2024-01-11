package main

import "fmt"

func main() {
	Array := []int{9, 3, 2, 5, 6, 3, 4, 72}
	QuickSort(Array,0,len(Array)-1)
	fmt.Print("Sorted array: ")
	for _, v := range Array {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func QuickSort(arr []int, low int, high int) []int {
	if low < high {
		pivot := Partition(arr, low, high)
		QuickSort(arr, low, pivot)
		QuickSort(arr, pivot+1, high)
	}
	return arr
}

func Partition(arr []int, low, high int) int {

	pivot := arr[low]
	leftIndex := low

	for i := low + 1; i <= high; i++ {
		if arr[i] < pivot {
			leftIndex++
			arr[i], arr[leftIndex] = arr[leftIndex], arr[i]
		}
	}
	arr[leftIndex], arr[low] = arr[low], arr[leftIndex]

	return leftIndex
}
