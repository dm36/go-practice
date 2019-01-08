package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10))
	fmt.Println(
		pow(3, 3, 20))
}

// Notes-
// - Variables declared inside an if are also available in the else
// statement
// - pow (3, 2, 10) => 3^2 = 9 < 10 so returns 9 (v)
// - pow(3, 3, 20) => 3^3 = 27 > 20 so print 27 >= 20 and return 20 (lim)
