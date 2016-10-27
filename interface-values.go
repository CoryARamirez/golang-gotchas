package main

import "fmt"

func main() {

	// START OMIT

	var (
		data *byte       // data is a pointer to a slice, data IS something
		in   interface{} // in is an interface, it points nowhere and it's nothing
	)

	// so, is "data" nil?
	fmt.Println(data, data == nil) // prints: <nil> true // HL

	// is "in" nil too?
	fmt.Println(in, in == nil) //prints: <nil> true // HL

	// Let's assign data (which is nil) to an empty interface
	in = data

	// And let's verify if in is still nil
	fmt.Println(in, in == nil) //prints: <nil> false // HL

	// END OMIT
}
