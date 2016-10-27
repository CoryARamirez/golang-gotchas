package main

import (
	"fmt"
	"time"
)

type files struct {
	name string
}

func (p *files) close() {
	fmt.Println("Closing:", p.name)
}

func main() {
	// START OMIT

	data := []files{{"one.pdf"}, {"two.pdf"}, {"three.pdf"}}

	for _, v := range data {
		go v.close()
	}

	// END OMIT

	time.Sleep(1 * time.Second)
}
