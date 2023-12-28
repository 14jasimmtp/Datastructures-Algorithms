package main

import "fmt"

func main() {
	slice := []string{"jasim", "sreehari", "ashkar"}
	for i, str := range slice {
		strRune := []rune(str)
		for i := range strRune {
			strRune[i]=strRune[i]+2
		}
		slice[i] = string(strRune)
	}
	fmt.Println(slice)
}
