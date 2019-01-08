package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

// Notes
// - Var declaration can include initializers- one per variable (1, 2 are the
// initializers in this case)
// - If initializer is present type can be omitted, variable takes type of the
// intializer- so c, python, java take the type of true, false, "no!"
