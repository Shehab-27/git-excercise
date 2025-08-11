//✳️ Task:
//Write a function called UniqueInts(nums []int) []int that returns a new slice containing only the unique integers from the input slice, in the order of their first appearance.

//✳️ Example:
//input := []int{1, 2, 2, 3, 4, 1, 5}
//result := UniqueInts(input)
//fmt.Println(result)
//✅ Expected Output:
//[]int{1, 2, 3, 4, 5}

package main

import "fmt"

func UniqueInts(nums []int) []int {
	var s []int
	m := make(map[int]bool)
	for _, val := range nums {
		if m[val] == false {
			m[val] = true
			s = append(s, val)
		}
	}
	return s
}

func main() {
	input := []int{1, 2, 2, 3, 4, 1, 5}
	result := UniqueInts(input)
	fmt.Println(result)
}
