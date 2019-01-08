package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// Notes-
// - Channels can be buffered, provide buffer length as the second argument to
//   make to initialize a buffered channel
// - Sends to a buffered channel block when the buffer is full, receives block
//   when the buffer is empty
// - So the send on line 9 is blocking because the buffer is full
