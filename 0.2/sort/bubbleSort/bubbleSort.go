package main

import "fmt"

func main() {
	Array := []int{9, 3, 2, 5, 6, 3, 4, 72}
	bubbleSort(Array)
	fmt.Print("Sorted array: ")
	for _, v := range Array {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
func bubbleSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
