// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 223.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

// So when you type Hello?- the main goroutine sends the Hello
// to the echo server over the connection request

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Connection to standard out, created goroutine, server response
	// to standard out

	go mustCopy(os.Stdout, conn)

	//Terminal input to the server (main goroutine)
	mustCopy(conn, os.Stdin)
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
