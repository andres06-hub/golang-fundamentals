package main

import "fmt"

func main() {
	// Slicen are mutable
	numbers := []int{1, 22, 73, 6, 5}
	fmt.Println(numbers, len(numbers)) // output: [1 2 3 4 5] 5

	// Add data at the end
	numbers = append(numbers, 8)
	numbers = append(numbers, 101)
	fmt.Println(numbers, len(numbers)) // output: [1 2 3 4 5] 5

	// Slice as a reference
	s3 := []int{4, 5, 9}
	s3_cpy := s3
	s3_cpy[0] = 1
	fmt.Println(s3, s3_cpy) // output: [1 5 9] [1 5 9]

	//Capacity
	months := []string{"Jan", "Feb", "Mar", "Apr"}
	//len() -> array Length
	// cap() -> array capacity
	// %p -> memory reference
	fmt.Printf("length: %d, Capacity: %d, %p\n", len(months), cap(months), months) // output: length: 4, Capacity: 4, 0xc000106040

	//Function make()
	// Params: type, length, capacity
	example := make([]int, 3, 3)
	example[0] = 10
	fmt.Println(example, len(example), cap(example)) // output: [10 0 0] 3 3
	example = append(example, 100)
	fmt.Println(example, len(example), cap(example)) // output: [10 0 0 100] 4 6
}
