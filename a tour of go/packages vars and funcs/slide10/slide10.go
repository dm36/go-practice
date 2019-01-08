package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

// Notes
// - := short assignment declaration can be used in place of a var declaration
// with implicit type
// - outside of a function, every statement begins with a keyword (var, func
// and so on) and the := construct is not available (outside of a func)
