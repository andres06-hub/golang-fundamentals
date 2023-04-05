package main

import "fmt"

func main() {
	// Arrays
	var a1 [4]int
	a1[0] = 2
	a1[3] = 5

	fmt.Println("a1: ", a1) //output: a1:  [2 0 0 5]

	// Define the length
	a2 := [3]string{"example1", "example2", "example3"}
	fmt.Println("a2: ", a2) //output: a2:  [example1 example2 example3]

	// Does not define the length
	a3 := [...]int{1, 2, 3, 4}
	fmt.Println("a3: ", a3) //output: a3:  [1 2 3 4]

	//defined indexes
	countries := [...]string{
		0: "USA",
		1: "COL",
		2: "BRA",
		3: "ESP",
	}
	fmt.Println(countries)      //output: [USA COL BRA ESP]
	fmt.Println(countries[1:4]) //output: [COL BRA ESP]
}
