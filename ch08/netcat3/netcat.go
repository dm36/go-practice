// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	//Listen
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	//Make a channel- channels are used to communicate between goroutines
	done := make(chan struct{})

	// Creates a goroutine which calls func- Copy copies contents of the
	// connection to standard out and then signals the main go routine that
	// it's completed copying

	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()

	// Sends the hello over to the reverb server and then closes the
	// connection
	mustCopy(conn, os.Stdin)
	conn.Close()

	<-done // wait for background goroutine to finish
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
