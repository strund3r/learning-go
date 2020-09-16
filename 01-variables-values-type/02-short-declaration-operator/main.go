package main

import "fmt"

func main() {
	// short declaration only works inside a func body!
	// := means the variable is being declared(:) and assigned(=) at the same time.
	x := 42
	fmt.Println(x)
	x = 007
	fmt.Println(x)
	name := "James Bond"
	fmt.Println(name)
	y := 100 + 24
	fmt.Println(y)
}
