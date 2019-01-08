package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// Notes-
// - Defer defers execution of the program until
// surrounding function returns.
// - Defer's arguments are evaluated immediately
//   but function call is not executed until surrounding
//   function returns
