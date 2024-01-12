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

func insertionSort(arr []int){
	for i:=1; i < len(arr); i++{
		current:=arr[i]

		j:=i-1
		for j>=0 && arr[j] > current{
			arr[j+1],arr[j]=arr[j],arr[j+1]
			j--
		}
		arr[j+1] = current
	}
}
