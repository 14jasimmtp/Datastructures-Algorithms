package main

import "fmt"

func reverseArray(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	reverseArray(slice)
	fmt.Println(slice)
}
