package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	x := Vertex{1, 2}
	p := &x
	p.X = 100
	fmt.Println(x)
	(*p).X = 200
	fmt.Println(x)
}

// Notes-
// - Struct fields can be accessed through a struct pointer
// - Can either do (*p).X or p.X
