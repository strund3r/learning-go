package main

import "fmt"

// DECLARE the variable "y"
// ASSIGN the value 43
// declare & assign = initialization
var y = 43

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
//
// Each element of such a variable or value is set to the zero value for its type:
// false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	// Use VAR when you need to declare a variable outside a func body!
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	foo()
}

func foo() {
	fmt.Println("again:", y)
}
