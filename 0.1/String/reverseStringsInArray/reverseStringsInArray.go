package main

import "fmt"

func main() {
	slice := []string{"jasim", "sreehari", "ashkar"}

	for i, str := range slice {
		strRune := []rune(str)
		left, right := 0, len(strRune)-1

		for left < right {
			strRune[left], strRune[right] = strRune[right], strRune[left]
			left++
			right--
		}

		slice[i] = string(strRune)

	}
	fmt.Println(slice)
}