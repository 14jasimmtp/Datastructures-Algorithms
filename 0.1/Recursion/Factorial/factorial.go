package main

import "fmt"

func factorial(val int) int{
	if val <= 1{
		return 1
	}

	return val*factorial(val-1)
	
}
func main() {
	display := factorial(3)
	fmt.Println(display)
}
