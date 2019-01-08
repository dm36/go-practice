package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Notes-
// - First parameter s of append is a slice of type T, and the rest
//   are T values to append to the slice
// - Resulting value of append is a slice containing all the elements
//   of the original slice plus the provided values
// - If backing array of s is too small to fit the given values, a bigger array
//   will be allocated, returned slice will point to the newly allocated array
// - Check out https://blog.golang.org/go-slices-usage-and-internals and
//   https://golang.org/pkg/builtin/#append
