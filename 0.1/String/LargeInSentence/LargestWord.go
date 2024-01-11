package main

import "fmt"


func findLargestWord(sentence string) string {
	var largestWord string
	var currentWord string
	var maxLength int

	for _, char := range sentence {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			currentWord += string(char)
		} else {
			if len(currentWord) > maxLength {
				maxLength = len(currentWord)
				largestWord = currentWord
			}
			currentWord = ""
		}
	}

	// Check for the last word in case the sentence doesn't end with a non-alphabetic character
	if len(currentWord) > maxLength {
		largestWord = currentWord
	}

	return largestWord
}

func main() {
	// Example usage
	sentence := "The quick brown fox jumpse over the lazy dog"
	largestWord := findLargestWord(sentence)

	fmt.Printf("The largest word in the sentence is: %s\n", largestWord)
}
