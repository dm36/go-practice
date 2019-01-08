package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

// Notes
// - Constants are declared like variables but with the const keyword
// - Can be character, string, boolean or numeric values, cannot be declared
// with := syntax
