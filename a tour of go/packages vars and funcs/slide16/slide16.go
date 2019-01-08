package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// 2 * 10 + 1
	fmt.Println(needInt(Small))
	// 2 * 0.1
	fmt.Println(needFloat(Small))
	// 1 < 100 * 0.1
	fmt.Println(needFloat(Big))
	fmt.Println(needInt(Big))
}

// Notes
// - Numeric constants are high-precision
// - Untyped constant takes the type needed by its context
// - Int can store at max a 64 bit integer and sometimes less
// which is why the needInt(Big) overflows
