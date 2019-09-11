package main

import "fmt"

// Var can be declared @ a global/package level
// This allows the variable to be used locally be any entity
var y string = "How are you?"

func main() {
	// The short declaration can only be accessed via its locally scope, and can not be hoisted
	x := "Hello"
	y = "world"
	n, _ := fmt.Println(x, y)
	foo()
	fmt.Println("something more.")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("This is an even number", i)
		}
	}
	fmt.Println(n)
}

func foo() {
	fmt.Println("I'm in foo")
}
