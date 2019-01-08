package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)
}

// Notes-
// - Array fixed-size
// - Slice- dynamically sized, flexible view into the elements
// - Type []T is a slice with elements of type T
// - Slice is formed by specifying two indices, low and high
//   separated by a colon
// - Includes first element but excludes last element
