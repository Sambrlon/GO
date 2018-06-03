package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 123)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 10, 20, 30, 124, 0)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}
