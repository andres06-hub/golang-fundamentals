package main

import "fmt"

func main() {

	m := map[string]int{"three": 3, "four": 4}

	// Like if, for doesn't use parens either.
	// Variables declared in for and if are local to their scope.
	for x := 0; x < 3; x++ { // ++ is a statement.
		fmt.Println("iteration", x)
		/*Output:
		iteration 0
		iteration 1
		iteration 2
		*/
	}

	// For is the only loop statement in Go, but it has alternative forms.
	for { // Infinite loop.
		break    // Infinite loop.
		continue // Infinite loop.
	}

	// You can use range to iterate over an array, a slice, a string, a map, or a channel.
	// range returns one (channel) or two values (array, slice, string and map).
	for key, value := range m {
		// for each pair in the map, print key and value
		fmt.Printf("key=%s, value=%d\n", key, value)
		/*Output
		key=three, value=3
		key=four, value=4
		*/
	}
	// If you only need the value, use the underscore as the key
	for _, name := range []string{"Bob", "Bill", "Joe"} {
		fmt.Printf("Hello, %s\n", name)
		/*output
		Hello, Bob
		Hello, Bill
		Hello, Joe
		*/
	}
}
