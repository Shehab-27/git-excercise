package main

import (
	"fmt"
	"strings"
)

func wordcount(s string) map[string]int {
	m := make(map[string]int)

	ipt2 := strings.ToLower(s)
	input := strings.Fields(ipt2)
	for _, word := range input {
		ipt := strings.Trim(word, "!,.")
		m[ipt]++
	}
	return m
}
func main() {
	input := "Go go go! Time to Go, now."
	result := wordcount(input)
	fmt.Println(result)
}
