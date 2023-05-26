package main

import (
	"fmt"
	"sort"
)

func findSmallestMissingPositive(nums []int) int {
	sort.Ints(nums)
	smallestMissing := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			continue
		}

		if nums[i] == smallestMissing {
			smallestMissing++
		} else if nums[i] > smallestMissing {
			break
		}
	}

	return smallestMissing
}

func main() {
	// Contoh 1
	input1 := []int{5, 2, 8, 4, 3, 10}
	output1 := findSmallestMissingPositive(input1)
	fmt.Printf("Input: %v\nOutput: %d\n\n", input1, output1)

	// Contoh 2
	input2 := []int{2, 3, 4, 6}
	output2 := findSmallestMissingPositive(input2)
	fmt.Printf("Input: %v\nOutput: %d\n\n", input2, output2)

	// Contoh 3
	input3 := []int{8, 6, 7, 12}
	output3 := findSmallestMissingPositive(input3)
	fmt.Printf("Input: %v\nOutput: %d\n\n", input3, output3)
}
