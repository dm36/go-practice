package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	var x = 10
	var y = 20
	fmt.Printf("The sum of %d and %d is %d", x, y, add(10, 20))
}

// Notes:
// - Function can take zero or more arguments
// - Type comes after the variable name
