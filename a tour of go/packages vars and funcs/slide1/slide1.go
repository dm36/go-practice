package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My fav number is: ", rand.Intn(10))
}

// Notes:
// - programs start running in package main
// - program is using packages with import paths "fmt" and "math/random
// - package name is the last element of the import path- i.e. /
//   rand in this case
