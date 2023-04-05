package main

import (
	"fmt"
)

func main() {

	var x float32
	x = 3.1416
	// z := "Name"
	g := 'Σ'

	// %s-> Refers to the formatting and printing of a string.
	fmt.Printf("Hello %s \n", "world")
	// %d -> Refers to the formatting and printing of a number.
	fmt.Printf("Port: %d \n", 3000)
	// %v -> any type of data
	fmt.Printf("any: %v , any: %v\n", "Yes", x)
	//
	msg := fmt.Sprintf("... %s, %d", "Example message", 8080)
	fmt.Println(msg)

	// know the type of data
	fmt.Printf("TypeOf: %T\n", g)
}
