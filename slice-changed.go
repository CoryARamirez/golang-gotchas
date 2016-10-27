package main

import (
	"fmt"
)

// START OMIT

func SubtractOneFromLength(slice []int) []int {
	slice = slice[:len(slice)-1] // reassign the slice, which should change it // HL
	return slice
}

func main() {
	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println("len(original) =", len(original), "\n") // HL

	modified := SubtractOneFromLength(original)
	fmt.Println("len(original) =", len(original)) // HL
	fmt.Println("len(modified) =", len(modified)) // HL
}

// END OMIT
