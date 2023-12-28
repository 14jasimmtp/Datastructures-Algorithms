package main

import "fmt"

func findDuplicates(arr []int) bool {
	seen := make(map[int]bool)
	for _, v := range arr {
		if seen[v] {
			return true
		}
		seen[v] = true
	}
	return false
}

func main() {
	slice := []int{1, 1, 4, 3, 4, 2, 7, 3}
	fmt.Println(findDuplicates(slice))
}
