package main

import "fmt"

func main() {
	fmt.Println("Hello")
	foo()
	fmt.Println("something more.")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("This is an even number", i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo")
}
