package main

import "fmt"

func main() {
	// START OMIT

	a := []byte("Hell")
	b := append(a, []byte("Fire")...)
	c := append(a, []byte("Wate")...)
	fmt.Println(string(a), "-", string(b), "-", string(c))

	// END OMIT
}
