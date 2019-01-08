package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
	fmt.Printf("The value of math.Pi is: %g ", math.Pi)
}

// Notes
// - Exported if it begins with a capital letter, Pizza is an exported name
//  as is Pi, pizza and pi are not exported
// - When importing a package can only refer to exported names, unexported
//   names are not accessible from outside the package
