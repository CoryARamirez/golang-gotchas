package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	data := "é"
	fmt.Println("Russian E:", utf8.RuneCountInString(data))

	data = "é"
	fmt.Println("Spanish E:", utf8.RuneCountInString(data))
}
