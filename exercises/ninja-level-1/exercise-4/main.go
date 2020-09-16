// 1. Create your own type. Have the underlying type be an int.
//
// 2. Create a VARIABLE of your new TYPE with the IDENTIFIER "x" using the VAR keyword.
//
// 3. In func main:
//
// a. print out the value of the variable "x"
// b. print out the type of the variable "x"
// c. assign 42 to the VARIABLE "x" using the = OPERATOR
// d. print out the value of the variable "x"

package main

import "fmt"

type mytype int

var x mytype

func main() {
	fmt.Println("Value of x:", x)
	fmt.Printf("Type of x: %T\n", x)
	x = 42
	fmt.Println("Value of x:", x)
}
