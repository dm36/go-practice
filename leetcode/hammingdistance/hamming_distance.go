package main

import (
	"fmt"
	"strconv"
)

// Adds y zeroes to the beginning of string x and returns
// the zerosAdded string
func addZeroes(x string, y int) (zerosAdded string) {
	fmt.Printf("The length of the string is %d\n", len(x))
	zerosAdded = x
	for i := 0; i < y; i++ {
		zerosAdded = "0" + zerosAdded
	}
	fmt.Printf("Zeroes added is %s\n", zerosAdded)
	return
}

func hammingDistance(x int, y int) (hamming int) {
	// Convert x and y both into integers of type 64 FormatInt
	// only takes int64's
	temp1 := int64(x)
	temp2 := int64(y)

	// Convert these numbers to binary strings
	xString := strconv.FormatInt(temp1, 2)
	yString := strconv.FormatInt(temp2, 2)

	fmt.Println(xString)
	fmt.Println(yString)

	// Prepend zeroes to the smaller string (same number of zeroes as the
	// difference in the two string lengths)
	if len(xString) > len(yString) {
		yString = addZeroes(yString, len(xString)-len(yString))
	} else if len(yString) > len(xString) {
		xString = addZeroes(xString, len(yString)-len(xString))
	}

	// Iterate over the strings and increment a counter everytime they
	// differ
	for i := 0; i < len(xString); i++ {
		if xString[i] != yString[i] {
			hamming += 1
		}
	}

	return hamming
}

func main() {
	fmt.Printf("The hamming distance is %d", hammingDistance(14, 2))
}
