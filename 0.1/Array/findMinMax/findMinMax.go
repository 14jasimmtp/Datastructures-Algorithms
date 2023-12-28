package main

import "fmt"

func findMinMax(arr []int) (int, int) {
	if len(arr) == 0 {
		panic("Empty array")
	}
	max, min := arr[0], arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}

func main() {
	slice := []int{1, 2, 3, 2, 56, 36, 3}
	fmt.Println(slice)
	fmt.Println(findMinMax(slice))
}
