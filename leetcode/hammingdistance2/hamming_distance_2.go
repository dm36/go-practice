package main

import (
	"fmt"
	"strconv"
)

func hammingDistance(x int, y int) (z int) {
	// Take the xor of x and y and convert it to int64
	// FormatInt takes int64's
	temp := int64(x ^ y)
	var count = 0

	// Convert to binary
	var string = strconv.FormatInt(temp, 2)

	// Iterate over the string and increment the count
	// everytime a 1 is observed
	for i := 0; i < len(string); i++ {
		if string[i] == '1' {
			count++
		}
	}
	return count
}

func main() {
	x := 14
	y := 2
	fmt.Printf("The hamming distance of %d and %d is: %d", x, y, hammingDistance(x, y))
}

// Notes-
// 10 is 01010
// 20 is 10100
