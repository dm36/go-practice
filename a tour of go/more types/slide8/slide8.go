package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// Notes-
// - Slice does not store data just describes a section of the underlying
//  array
// - Changing the elements of a slice modifies the corresponding
//   elements of the underlying array
// - Other slices that share the same underlying array will see those changes
// - Reason that b[0]'s change propagates to a and b is because a and b
//   overlap on an index
