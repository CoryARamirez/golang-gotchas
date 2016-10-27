package main

import "fmt"

// START OMIT

func main() {
	var slice []string

	if slice != nil {
		for _, str := range slice {
			fmt.Println("String:", str)
		}
	}
}

// END OMIT
