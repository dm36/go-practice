package main

import "fmt"

// Struct called Vertex containing two int's
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	// Store vertex to variable x and print out
	// its fields
	x := Vertex{1, 2}
	fmt.Println(x.X)
	fmt.Println(x.Y)
}

// Notes
// - Struct is a collection of fields
