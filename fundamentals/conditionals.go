package main

import "fmt"

func main() {
	// If statements require brace brackets, and do not require parentheses.
	if true {
		fmt.Println("told ya")
	}
	// Formatting is standardized by the command line command "go fmt".
	if false {
		// Pout.
	} else {
		// Gloat.
	}
	// Use switch in preference to chained if statements.
	x := 42.0
	switch x {
	case 0:
		fmt.Println(100)
	case 1, 2: // Can have multiple matches on one case
		fmt.Println(200)
	case 42:
		fmt.Println(300)
		// Cases don't "fall through".
		/*
			There is a `fallthrough` keyword however, see:
			https://github.com/golang/go/wiki/Switch#fall-through
		*/
	case 43:
		fmt.Println(400)
		// Unreached.
	default:
		fmt.Println("Default")
		// Default case is optional.
	}
}
