package main

import "fmt"

var y = 42

// DECLARED the VARIABLE with the IDENTIFIER "z" is of TYPE string
// and I ASSIGN the VALUE "Shaken, not stirred"
var z string = "Shaken, not stirred"

var a string = `James said 
"Shaken, 
not stirred"`

// this is a STATIC programing language
// not a DYNAMIC programing language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE

func main() {
	fmt.Println(y)
	fmt.Printf("Value: %d, Type: %T\n", y, y)
	fmt.Println(z)
	fmt.Printf("Value: %v, Type: %T\n", z, z)
	fmt.Println(a)
}
