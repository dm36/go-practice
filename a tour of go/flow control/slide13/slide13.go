package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 10; i > 0; i-- {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// Notes -
// - Deferred function calls are pushed onto a stack
// - When a function returns- deferred calls are executed in
//   LIFO order
// - https://blog.golang.org/defer-panic-and-recover
