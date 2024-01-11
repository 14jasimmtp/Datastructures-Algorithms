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

// package main

// import "fmt"

// func replaceAlphabet(input string, n int) string {
// 	result := ""
// 	for _, char := range input {
// 		if isAlphabet(char) {
// 			base := 'a'
// 			if char >= 'A' && char <= 'Z' {
// 				base = 'A'
// 			}
// 			newChar := base + (char-base+rune(n))%26
// 			result += string(newChar)
// 		} else {
// 			result += string(char)
// 		}
// 	}
// 	return result
// }

// func isAlphabet(char rune) bool {
// 	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
// }

// func main() {
// 	inputString := "Hello, zzzz!"
// 	n := 3
// 	result := replaceAlphabet(inputString, n)
// 	fmt.Println(result)
// }
