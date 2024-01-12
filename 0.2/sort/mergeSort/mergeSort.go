package main

import "fmt"

func main() {
	Array := []int{9, 3, 2, 5, 6, 3, 4, 72}
	
	fmt.Print("Sorted array: ")
	for _, v := range mergeSort(Array) {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result

}

// func mergeSort(arr[]int) {
// 	if len(arr) < 2{
// 		return 
// 	}

// 	mid:=len(arr)/2
// 	mergeSort(arr[:mid])
// 	mergeSort(arr[mid:])
// 	merge(arr[:mid],arr[mid:])
// }

// func merge(left,right []int) {
// 	var result []int
// 	i,j:= 0,0
// 	for len(left) > i && len(right) > j{
// 		if left[i] < right[j] {
// 			result = append(result,left[i])
// 			i++
// 		}else{
// 			result = append(result,right[j])
// 			j++
// 		}
// 	}
// 	result = append(result,left[i:]...)
// 	result = append(result,right[j:]...)
// }