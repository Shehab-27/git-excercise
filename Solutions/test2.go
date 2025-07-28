package main

import "fmt"

func UniqueInts(nums []int) []int {
	//init map
	u := make(map[int]bool)
	result := []int{}
	//remove repeated nums
	for _, n := range nums {
		if !u[n] {
			u[n] = true
			result = append(result, n)
		}
	}
	return result
}

func main() {
	input := []int{1, 2, 2, 3, 4, 1, 5}
	result := UniqueInts(input)
	fmt.Println(result)
}
