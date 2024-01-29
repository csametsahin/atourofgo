package main

import "fmt"

func main() {
	// initializes the function but calls it after surronding function return
	defer fmt.Println("world")
	// LİFO
	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
