package main

import "fmt"

func main() {
	v := 0.5 + 20i
	fmt.Printf("v is of type %T\n", v)
}

// Notes-
// - When you declare a variable without an explicit type (using either := or
//  var = expression syntax), variable's type is inferred from the value
//  on the right side
// - When the right hand side of the declaration is typed- the new variable is
//   of the same type
//   var i int
//   j:= i // j is also an int
// - Right hand contains an untyped numeric constant, the new variable may be
//   an int, float64 or complex128 based on precision of the constant
