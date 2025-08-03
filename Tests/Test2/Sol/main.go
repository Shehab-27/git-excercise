package main

import "fmt"

// ✳️ Task:
// returns a new slice containing only the unique integers from the input slice, in the order of their first appearance.

// ✳️ Example:

// ✅ Expected Output:
// []int{1, 2, 3, 4, 5}

func UniqueInts(nums []int) []int {
	freq := make(map[int]bool)
	var arr []int

	for _, num := range nums {
		if !freq[num] {
			freq[num] = true
			arr = append(arr, num)

		}

	}
	return arr

}
func main() {

	input := []int{1, 2, 2, 3, 4, 1, 5, 5, 6, 6, 7, 9}
	result := UniqueInts(input)
	fmt.Println(result)

}
