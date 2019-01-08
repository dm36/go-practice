package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	x := Vertex{5, 10}
	x.X = 20
	fmt.Println(x)
	fmt.Println(x.X)
	fmt.Println(x.Y)
}

// Notes-
// - Can acess struct fields with a dot
