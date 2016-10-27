package main

import (
	"fmt"
)

func main() {
	// START OMIT

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, n := range numbers {
		defer func() {
			fmt.Println("Number:", n)
		}()
	}

	// END OMIT
}
