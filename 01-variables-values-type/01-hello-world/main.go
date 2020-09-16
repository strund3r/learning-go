package main

import "fmt"

func foo() {
	fmt.Println("I'm in foo.")
}

func bar() {
	fmt.Println("and then we exited.")
}

func main() {
	fmt.Println("Hello, world.")

	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

// control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional
