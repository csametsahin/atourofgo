// package main

// import "fmt"

// var dotnet, java, python bool

//	func main() {
//		var i int
//		fmt.Println(i, dotnet, java, python)
//	}
package main

import "fmt"

var x, y int = 1, 2

func main() {
	// did not user var used := variable initializer
	c, python, java := true, false, "never"
	fmt.Println(c, python, java, x, y)
}
