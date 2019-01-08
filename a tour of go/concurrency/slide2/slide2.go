package main

import "fmt"

// Arguments are an array of integers s and a channel of integers c
func sum(s []int, c chan int) {
	// Iterate over the elements of the array s and compute the sum and then
	// send the sum over the channel
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// Make a channel of integers
	c := make(chan int)

	// Go routine for the left side of the array and the right side
	// off the array- the sum function will send the sum through the channel
	// and then the main routine will pull the two off the channel into x and y
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

// Notes-
// - Channels are a type conduit through which you can send and receive values
//   with the channel operator
// - ch <- v // Send v to the channel ch
//   v := <-ch // Receive from ch and assign value to v
// - Like maps and slices channels must be created before use (ch := make(chan int))
// - Send and receives block until other side is ready, allows goroutines
//   to sycnhronize without explicit locks or condition variables
// - Sums the numbers in a slice, distributes the work between two goroutines, once
//   goroutines have completed their computation calculates the final result
