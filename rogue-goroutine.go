package main

import "fmt"

func main() {

	// START OMIT

	done := false

	go func() {
		done = true
	}()

	for !done {
		// simulate I/O blocking
	}

	fmt.Println("done!")

	// END OMIT
}
