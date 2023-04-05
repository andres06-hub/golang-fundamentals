package main

import (
	"fmt"
)

func main() {
	fmt.Println(learnMultiple(2, 4))
	greet("Hello", "World")
	fmt.Println(sum(5, 3))
}

// function with parameters
func greet(msg string, name string) {
	fmt.Println(msg, name)
}

// Fucntion with return type
func sum(x, y int) int {
	return x + y
}

/*
Functions can have parameters and (multiple!) return values.
Here `x`, `y` are the arguments and `sum`, `prod` is the signature (what's returned).
Note that `x` and `sum` receive the type `int`.
*/
func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y // Return two values.
}
