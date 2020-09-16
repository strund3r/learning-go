// 1. At the package level scope,
// using the "var" keyword, create a VARIABLE with the IDENTIFIER "y".
// The variable should be of the UNDERLYING TYPE of your custom TYPE "x".
//
// 2. In func main:
//
// a. print out the value of the variable "x"
// b. print out the type of the variable "x"
// c. assign your OWN VALUE to the VARIABLE "x" using the = OPERATOR
// d. print out the value of the variable "x"
// e. use CONVERSION to convert the type of the VALUE stored in "x" to the UNDERLYING TYPE
// f. then use the SHORT DECLARATION OPERATOR to ASSIGN that value to "y"
// g. print out the value stored in "y"
// h. print out the type of "y"

package main

import "fmt"

type mytype int

var x mytype
var y int

func main() {
	fmt.Println("Value of x:", x)
	fmt.Printf("Type of x: %T\n", x)
	x = 21
	fmt.Println("Value of x:", x)

	y = int(x)
	fmt.Println("Value of y:", y)
	fmt.Printf("Type of y: %T\n", y)

}
