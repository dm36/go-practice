package main

import "fmt"

func sum(x int, y int) (z int) {
	z = x + y
	return
}

func main() {
	x := 10
	y := 5
	fmt.Printf("%d", sum(x, y))
}

// Notes-
// - Return statement without arguments returns the named
// returned values- naked return (z int)
// - Should only be used in short functions
