package main

import "fmt"

func printBilanganCacah(N int) {
	count := 0
	i := 1

	for count < N {
		if i%3 == 0 && i%7 == 0 {
			fmt.Print("Z, ")
		} else if i&3 == 0 || i&7 == 0 {
			fmt.Printf("%d, ", i)
		}
		count++
		i++
	}
}

func main() {
	N := 13
	fmt.Printf("output: ")
	printBilanganCacah(N)
}
