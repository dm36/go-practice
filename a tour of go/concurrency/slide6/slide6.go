package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	// My understanding is we hit a tick every 100 milliseconds,
	// a boom every 500 milliseconds and otherwise we're hitting the .A
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("     .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}

// Notes-
// - The default case in a select is run if no other
//   case is ready
// - Use a default case to try and send or receive without
//   blocking
// select {
// case i := <-c:
//     // use i
// default:
// receiving from c would block
// }
//
//
