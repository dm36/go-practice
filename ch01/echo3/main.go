// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

// All in one line, join the command line arguments
// with a white space and then print them out
//!+
func main() {
	fmt.Println(strings.Join(os.Args[1:], ", "))
}

//!-
