// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 12.

//!+

// Dup3 prints the count and text of lines that
// appear more than once in the named input files.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
// Read the entire contents of a file in one big gulp
// split into lines all at once, then process the lines

// Reads named files not standard input, moved counting back
// to main
func main() {
	// Same counts mapping
	counts := make(map[string]int)

	// Iterate over the files, read the file in a single gulp
	// split in based on newlines and build the counts mapping
	// accordingly
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	// Check to see if any of the counts are above 1 and report those
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-
