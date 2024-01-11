package main

import "fmt"

func main() {
	Array := []int{9, 3, 2, 5, 6, 3, 4, 72}
	insertionSort(Array)
	fmt.Print("Sorted array: ")
	for _, v := range Array {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
func insertionSort(arr []int) {
	for i := 1; i < len(arr)-1; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
