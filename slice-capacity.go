package main

import "fmt"

// START OMIT

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // HL

func main() {
	var evenNumbers []int // HL

	for _, n := range numbers {
		if n%2 == 0 {
			evenNumbers = append(evenNumbers, n) // HL
		}
	}

	fmt.Printf("%#v", evenNumbers)
}

// END OMIT
