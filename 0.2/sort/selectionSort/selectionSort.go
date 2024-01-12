package main

import "fmt"

func main() {
	Array := []int{9, 3, 2, 5, 6, 3, 4, 72}
	selectionSort(Array)
	fmt.Print("Sorted array: ")
	for _, v := range Array {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func selectionSort(arr []int) {
	for i:=0;i<len(arr);i++{
		minIndex:=i
		for j:=i+1 ; j<len(arr);j++{
			if arr[j] < arr[minIndex]{
				minIndex = j
			}
		}
		arr[i],arr[minIndex] = arr[minIndex],arr[i]
	}
}
