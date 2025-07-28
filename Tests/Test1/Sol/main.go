package main

import (
	"fmt"
	"strings"
)

// COUNT REPEATED WORDS
func wordCount(s string) map[string]int {

	// init map
	result := make(map[string]int)

	// trimmable punctuations
	punctuations := "!,.?"

	// convert text to lower (case sensetivity)
	lowerText := strings.ToLower(s)

	// split string
	arr := strings.Fields(lowerText)

	// trim punctuations and check map for key repetition
	for _, word := range arr {
		newWord := strings.Trim(word, punctuations)
		result[newWord]++
	}

	return result
}

func main() {

	text := "Go, go go! time to go now."
	result := wordCount(text)
	fmt.Println(result)

}
