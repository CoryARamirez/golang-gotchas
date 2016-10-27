package main

import (
	"bytes"
	"fmt"
	"strings"
)

var data = []byte("this is a JSON body from HTTP req!")

func main() {
	// START OMIT

	// Type conversion with string(data)
	if strings.Contains(string(data), "body") { // HL
		fmt.Println("[str] The http body does contain the word 'body'!")
	} else {
		fmt.Println("[str] Word 'body' not found in http body.")
	}

	// We type-convert the data we know, not the unknown one
	if bytes.Contains(data, []byte("body")) { // HL
		fmt.Println("[byte] The http body does contain the word 'body'!")
	} else {
		fmt.Println("[byte] Word 'body' not found in http body.")
	}

	// END OMIT
}
