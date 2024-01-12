package main

import "fmt"

func main() {
	Array := []int{9, 3, 2, 5, 6, 3, 4, 72}
	res:=quickSort(Array)
	fmt.Print("Sorted array: ")
	for _, v := range res {
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


func quickSort(arr []int) []int{

	if len(arr) < 2{
		return arr
	}

	pivot:=arr[(len(arr)-1)/2]

	var left,right []int

	for i:=1;i<len(arr);i++{
		
		if arr[i] < pivot{
			left = append(left,arr[i])
		} else{
			right = append(right,arr[i])
		}
	}

	return append(append(quickSort(left),pivot), quickSort(right)...)
}