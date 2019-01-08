package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer (42)
	*p = 21         // set i through the pointer
	fmt.Println(i)  // i now has the value 21

	p = &j         // point to j
	*p = *p / 37   // divide j by 37
	fmt.Println(j) // print j
}

// Notes-
// - Type *t is a pointer to a T value, zero value is nil
// - Declared as such var p *int
// - & generates a pointer to its operand
// - i := 42, p = &i, int i with value 42, pointer to int
// - * operator denotes the pointer's underlying value,
//   dereferencing or indirecting
// - Go has no pointer arithmetic- which C does
