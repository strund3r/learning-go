package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("Value: %d, Type: %T\n", a, a)
	fmt.Println(b)
	fmt.Printf("Value: %v, Type: %T\n", b, b)
	a = int(b)
	fmt.Println(a)
	fmt.Printf("Value: %d, Type: %T\n", a, a)
}
