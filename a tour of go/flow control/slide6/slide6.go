package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

// Notes-
// - If can start with a short statement to execute before the condition
// (if v:= math.Pow(x, n))
// - Variables declared by if statement are only in scope until the end of the
// if (v can not be used outside of the if statement)
