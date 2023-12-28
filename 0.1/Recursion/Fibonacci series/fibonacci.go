package main

import "fmt"


func fibonacci(val int) int{
	if val < 1{
		return val
	}
	return fibonacci(val-1) + fibonacci(val-2)
}
func main() {
	result := fibonacci(6)
	fmt.Println(result)
}
