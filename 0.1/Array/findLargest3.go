package main

import "fmt"

func FindLargestThree(slice []int) (int, int, int) {
	var largest, secondlargest, thdlargest int
	for _, n := range slice {
		if n > largest {
			thdlargest, secondlargest, largest = secondlargest, largest, n

		} else if n > secondlargest && n != largest {
			thdlargest, secondlargest = secondlargest, n
		} else if n > thdlargest && n != secondlargest {
			thdlargest = n
		}

	}
	return largest, secondlargest, thdlargest
}
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	lg, md, small := FindLargestThree(slice)
	fmt.Printf("%d is largest,%d is second largest ,%d is third largest", lg, md, small)
}
