package main

import "fmt"

func main() {
	slice := []string{"Jasim", "sreehari", "ashkar"}

	for i, str := range slice {
		if len(str) > 0 {
			firstchar := str[0]
			if 'a' <= firstchar && firstchar <= 'z' {
				firstchar = firstchar - ('a' - 'A')
			}
			slice[i] = string(firstchar) + str[1:]
		}
	}
	fmt.Println(slice)
}