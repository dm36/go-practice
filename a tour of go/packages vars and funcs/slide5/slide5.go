package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	var x = 30
	var y = 40
	fmt.Printf("The sum of %d and %d is: %d", x, y, add(x, y))
}

// Notes
// - Can shorten x int, y int to x, y int
