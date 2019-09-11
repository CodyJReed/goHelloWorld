package main

import "fmt"

var y string = "How are you?"

func main() {
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
