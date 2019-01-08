package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("Hello")
}

// Notes
// - Goroutine is a lightweight thread managed by the Go runtime
// - go f(x, y, z) starts a new goroutine running f(x, y, z)- evaluation
//   happens in the current goroutine and the execution of f happens
//   in the new goroutine.
// - Goroutines run in the same address space- so access to shared memory must
//   be synchronized
// - Sync package provides useful primitives
