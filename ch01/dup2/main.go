// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Counts maps strings to integers, files
	// is everything the user typed in the command
	// line aside from the command to execute the program
	counts := make(map[string]int)
	files := os.Args[1:]

	// If no files are provided on the command-line
	// call the countLines function passing standard in
	// and the counts mapping
	if len(files) == 0 {
		countLines(os.Stdin, counts)

	// Otherwise iterate over the files provided over command-line-
	// open each of the files and call the countLines function
	// passing the file pointer as the first argument
	// and the counts mapping as the second argument
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)

			// Closing the file
			f.Close()
		}
	}

	fmt.Println(counts)
	
	for line, n := range counts {
		fmt.Printf("%s %d\n", line, n)
	}

	// Go through our counts mapping which maps
	// lines to counts- and print out each of the lines
	// where the count is greater than 1
	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d\t%s\n", n, line)
	// 	}
	// }
}

// First argument passed is either standard in or
// the pointer to the file- the second argument is
// a mapping from strings to integers- i.e. lines
// to line counts
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
