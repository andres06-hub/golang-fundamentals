package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1
	fmt.Println(m)

	// add
	week := make(map[string]int)
	week["monday"] = 1
	week["tuesday"] = 2
	week["wednesday"] = 3
	week["thursday"] = 4
	week["friday"] = 5
	week["saturday"] = 6
	week["sunday"] = 7
	fmt.Println(week)

	// Delete
	delete(week, "sunday")
	fmt.Println(week)

	// new Map -> Array inside map
	students := make(map[string][]int)
	students["alex"] = []int{8, 5}
	students["andres"] = []int{6, 12}
	fmt.Println(students)              //output: map[alex:[8 5] andres:[6 12]]
	fmt.Println(students["andres"])    //output: [6 12]
	fmt.Println(students["andres"][0]) //output: 6
}
