package main

import "fmt"

// Add each element of the fibonacci sequence to the
// channel c
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10) // channel of 10 integers
	go fibonacci(cap(c), c) //call the fibonacci function as a goroutine and
	// pass the capacity of the channel as the first argument and the channel
	// as the second argument

	// This retrieves values from the channel and prints them
	for i := range c {
		fmt.Println(i)
	}
}

// Notes-
// - Sender can close a channel to indicate that no more values will be sent
//   (close(c))
// - Receivers can test whether a channel has been closed by assigning a second
//   parameter to the receive expression: after v, ok := <-ch
//   ok is false if there are no more values to receive and the channel is closed
// - loop for i := range c receives values from the channel repeatedly until it
//   is closed
// - sender should close a channel never the receiver, sending on a closed channel
//   will cause a panic
// - channels aren't like files- you don't usually need to close them, closing is only
//   needed when receiver must be told that there are no more values coming, to terminate
//   a range loop (to terminate the for i := range c)
