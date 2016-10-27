package main

import (
	"fmt"
)

// START OMIT

func main() {
	data := []string{"a", "b", "c"} // HL
	sl(data)
	fmt.Printf("%#v", data)
}

func sl(a []string) {
	if len(a) == 3 {
		a[0], a[1], a[2] = "d", "e", "f" // HL
	}
}

// END OMIT
