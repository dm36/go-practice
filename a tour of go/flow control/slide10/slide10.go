package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("Day after.")
	default:
		fmt.Println("Too far away.")
	}
}

// Notes-
// - Switch cases evaluate from top to bottom and stop when
// a case succeeds
// - switch i {
// case 0:
// case f():
// } does not stop if i == 0
// - Logic of the program-
//   we switch on Saturday- if Saturday is today + 0
//   it indicates today is Saturday, if Saturday is
//   today + 1 it indicates tomorrow is Saturday, so on a
//   and so forth
