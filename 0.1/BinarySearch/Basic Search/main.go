package main

import "fmt"

func BinarySearch(Array []int, value int) int {
	low := 0
	high := len(Array) - 1

	for low <= high {
		median := (low + high) / 2
		if Array[median] < value {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(Array) || Array[low] != value {
		return -2
	}

	return low + 1

}

func LinearSearch(Array []int, value int) int {
	for i, v := range Array {
		if v == value {
			return i
		}
	}
	return -2
}

func main() {
	Array := []int{1, 59, 45, 23, 56}
	SortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(LinearSearch(Array, 34) + 1)
	fmt.Println(BinarySearch(SortedArray, 7))
}
