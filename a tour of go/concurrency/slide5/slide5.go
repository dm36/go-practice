package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1

	// Pulls 0-10 off the channel
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// Two channels of integers
	c := make(chan int)
	quit := make(chan int)

	// Adding 0-10 to c and then 0 to quit is done by this
	// go routine
	go func() {
		// Think this for loop is adding the numbers 0-10 to the channel
		// (the one on slide 4 is retrieving numbers from the channel)
		// think the official terminology is send and receive from a
		// channel
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		// Put 0 on the quit channel after we add 0-10
		quit <- 0
	}()

	// Main goroutine which takes #s off the channel, computes the fibonacci
	// of that #- what's printing the fib #s- is it the fmt.Printnln(<-c)??
	fibonacci(c, quit)
}

// Notes-
// - Select statement lets a goroutine wait on multiple
//   communication operations
// - Select blocks until one of its cases can run- chooses one at random
//   if multiple are ready
