package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// Option 1: Make a local copy, which looks nasty and adds more complexity.
	for _, n := range numbers {
		nn := n // Avoid this! // HL
		go func() {
			fmt.Println("Number:", nn)
		}()
	}

	// Option 2: Pass the value, since Go passes by-value by default.
	for _, n := range numbers {
		go func(localnumber int) { // HL
			fmt.Println("Number:", n)
		}(n) // HL
	}

	// END OMIT

	time.Sleep(5 * time.Second)
}
