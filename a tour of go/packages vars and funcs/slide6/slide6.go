package main

import "fmt"

func swapStrings(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(swapStrings("dog", "cat"))
}

// Notes:
// - Function can return any number of results
// - Note the return part (string, string)- only need to specify
//   variable types not names
