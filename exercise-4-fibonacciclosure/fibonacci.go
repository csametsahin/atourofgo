package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// tamamen internetten aldım ama burada recuressivein farklı bir mantığı oluşmuş oluyor sanırım bir debuglarsam iyi olur
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		f := x
		x, y = y, f+y

		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
