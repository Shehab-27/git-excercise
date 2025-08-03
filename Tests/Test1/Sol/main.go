package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	new_s := strings.ToLower(s)
	arr := strings.Fields(new_s)

	freq := make(map[string]int)
	// counter := 0
	// for i:=0 ; i < len(arr); i++ {
	for _, word := range arr {
		batates := strings.Trim(word, ",.!?")
		freq[batates]++

		// if freq[word] == arr[i] {
		// 	freq[arr[i]] = counter + 1
		// }
	}
	return freq
}

func main() {

	input := "Go go go! Time to Go, now."

	result := WordCount(input)

	fmt.Println(result)
}
