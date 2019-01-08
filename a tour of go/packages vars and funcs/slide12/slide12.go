package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// Notes
// - Variables declared without an explicit initial value are given
// a zero value
// - Zero value is 0 for numeric types, false for boolean type and ""
// (empty string) for strings
