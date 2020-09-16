package main

import "fmt"

var y = 42

func main() {
	// Print the value of y
	fmt.Println(y)
	// Print the type of y
	fmt.Printf("%T\n", y)
	// Print binary form of y
	fmt.Printf("%b\n", y)
	// Print hexadecimal form of y
	fmt.Printf("%x\n", y)
	// Print hexadecimal form of y with 0x in front
	fmt.Printf("%#x\n", y)
}
