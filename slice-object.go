package main

// START OMIT

type slice struct {
	Length        int
	Capacity      int
	ZerothElement *byte
	Array         *[...]T // this is just a way to describe it // HL
}

// END OMIT
