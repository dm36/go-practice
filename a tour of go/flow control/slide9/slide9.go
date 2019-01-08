package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println(" OSX.")
	case "linux":
		fmt.Println(" Linux.")
	default:
		fmt.Printf(" %s.", os)
	}
}

// Notes-
// - Go runs on the selected case not the ones that follow
// - Break statement at the end of each case in other languages
//   is provided automatically in go
// - Go's switch need not be constants, and values involved
//   need not be integers
// - https://golang.org/pkg/runtime/
