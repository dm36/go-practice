package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// Notes:
// - T(v) converts the value v to the type V
// - Can either do:
// var i int = 42
// var f float64 = float64(i)
// var u uint = uint(f)
// - Or do:
// i := 42
// f := float64(i)
// u := uint(f)
// - Assignment between items of different type requires explicit conversions
