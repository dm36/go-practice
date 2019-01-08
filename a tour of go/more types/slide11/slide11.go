package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice and give it zero length
	s = s[:0]
	printSlice(s)

	// Extend its length
	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Notes-
// - Slice has a length and capacity, length is the # of elements it contains
// - Capacity is the # of elements in the underlying array- counting from
//   first element in the slice
// - Length and capacity of s can be obtained using the expression len(s)
//   and cap(s)
// - Can extend slice's length by re-slicing it, if it has sufficient capacity
