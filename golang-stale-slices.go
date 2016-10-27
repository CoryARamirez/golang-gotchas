package main

import "fmt"

func main() {
	// START OMIT

	// We create a first simple slice of 3 numbers
	slice1 := []int{1, 2, 3}
	fmt.Println("Length:", len(slice1), "Capacity:", cap(slice1), "Data:", slice1)

	// We create a second slice based on the first one
	slice2 := slice1[1:]
	fmt.Println("Length:", len(slice2), "Capacity:", cap(slice2), "Data:", slice2)

	// We change ALL the values in SLICE2 (not SLICE1)
	// to add 20! We didn't update SLICE1
	for i := range slice2 {
		slice2[i] += 20
	}

	// What will be the output?
	fmt.Println(slice1) // expected: [1 2 3]
	fmt.Println(slice2) // expected: [22 23]

	// END OMIT
}
