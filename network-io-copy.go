package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// START OMIT

	resp, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatal("Something happened:", err.Error())
	}

	defer resp.Body.Close()

	var buf bytes.Buffer     // HL
	io.Copy(&buf, resp.Body) // HL

	fmt.Println(buf.String())

	// END OMIT
}
