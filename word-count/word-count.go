package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Needles and pins " +
		"needles and pins " +
		"Sew me a sail " +
		"To catch me the Wind "

	// Account for case-insensitivity
	textLowercase := strings.ToLower(text)

	// Split text into words
	wordArray := strings.Fields(textLowercase)

	// Iterate over words to get count
	wordMap := map[string]int{}
	for _, word := range wordArray {
		//store the word and it's count to wordMap
		wordMap[word]++
	}

	// Iterate over words to report count
	for word, count := range wordMap {
		fmt.Printf("%s was used %d times\n", word, count)
	}
}
