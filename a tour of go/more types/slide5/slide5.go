package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	p  = &Vertex{1, 2} // pointer to a vertex struct containing a 1 and a 2
	v1 = Vertex{1, 2}  // x is 1 y is 2
	v2 = Vertex{X: 1}  // x is 1 y is 0
	v3 = Vertex{}      // x and y are both zero
)

func main() {
	fmt.Println(p, v1, v2, v3)
}

// Notes-
// - Struct literal denotes a newly allocated struct value by listing the values
// of its fields {1, 2}
// - List using a subset of fields by using the Name: syntax
// - & returns a pointer to the struct value
