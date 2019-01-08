package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// Notes-
// - Go only has one looping construct: the for loop
// - For loop has three components separated by semicolons
// - init statement, executed before the iteration, condition expression,
//   evaluated before every iteration, post statement- executed at the end of
//   every iteration
// - init statement will often be a short variable delaration, variables only
//   visible in the scope of the for statement
// - loop stops iterating once boolean condition evaluates to false
// - no parens surrounding the three components of the for statement and braces
//   always req (unlike C...)
