package main

import "fmt"

func main() {
	// START OMIT

	type Bank struct{ Name string }

	type Person struct {
		Name string
		Bank *Bank
	}

	var (
		bob  = Person{Name: "Bob"}
		lisa = Person{Name: "Lisa", Bank: &Bank{Name: "BoA"}}
	)

	fmt.Println("[👍]", lisa.Name, "bank is:", lisa.Bank.Name, "\n")
	fmt.Println("[👎] ", bob.Name, "bank is:", bob.Bank.Name)

	// END OMIT
}
