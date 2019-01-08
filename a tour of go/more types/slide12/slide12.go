package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

// Notes-
// - Zero value of a slice is nil
// - Nil slice has a length and capacity of 0 and no underlying array
