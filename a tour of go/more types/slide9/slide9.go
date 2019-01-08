package main

import "fmt"

func main() {
	// Slice of integers
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// Slice of booleans
	f := []bool{true, false, true, true, false, true}
	fmt.Println(f)

	// Slice of structs
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

// Notes
// - Slice literal is an array literal without the length
// - Array literal [3]bool{true, true, false}
// - []bool{true, true, false} creates the same array as above
//   then builds a slice that references it
