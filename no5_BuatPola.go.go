package main

import (
	"fmt"
)

func printPattern(N int) {
	if N%2 == 0 {
		fmt.Println("Harus bilangan ganjil")
		return
	}

	size := N
	if N < 0 {
		size = -N
	}

	pattern := make([][]rune, size)
	for i := range pattern {
		pattern[i] = make([]rune, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j || i+j == size-1 {
				pattern[i][j] = 'X'
			} else {
				pattern[i][j] = 'O'
			}
		}
	}

	for i := range pattern {
		for j := range pattern[i] {
			fmt.Printf("%c", pattern[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	N := 5
	printPattern(N)

	N = 3
	printPattern(N)

	N = 7
	printPattern(N)

	N = 2
	printPattern(N)
}
