// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Goes through the URL's passed on the command-line
	// make a get request to each- retrieve the repsonse from
	// resp.Body, and print it out
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
	}
}

//!-
