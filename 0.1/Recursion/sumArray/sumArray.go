package main

import "fmt"

func sumArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else {
		return arr[0] + sumArray(arr[1:])
	}
}

func main() {
	slice := []int{4, 2, 6, 13, 25, 23}
	fmt.Println(slice)
	fmt.Println(sumArray(slice))
}
