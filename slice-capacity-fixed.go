package main

import "fmt"

// START OMIT

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {
	// Instead of not declaring it explicitly, we do declare it now,
	// with the LENGTH of zero elements (since we don't know how many will be)
	// but with a CAPACITY of 10, since that's the maximum number we may have
	evenNumbers := make([]int, 0, len(numbers)) // HL

	for _, n := range numbers {
		if n%2 == 0 {
			evenNumbers = append(evenNumbers, n)
		}
	}

	fmt.Printf("%#v \n", evenNumbers)
	fmt.Printf("Size: %v", len(evenNumbers))
}

// END OMIT
