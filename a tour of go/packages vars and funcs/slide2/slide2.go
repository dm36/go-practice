package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The square root of the number 7 is: ", math.Sqrt(7))
}

//
// func main() {
// 	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
// }

// Notes:
// - Parenthesized, factored import statement- could also have done
//   import "fmt", import "math"
// - Factored, imported statement is good style
// - Note that fmt.Printf also exists, first argument is a string with format
//   specifiers, remaining arguments are the values for those format
//   specifiers
