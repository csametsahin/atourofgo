package main

import "fmt"

func main() {
	/*kinda like arrays but dont use bounds it creates an slice referance to this*/
	p := []int{1, 2, 3, 5, 7, 9}
	fmt.Println(p)
	r := []bool{false, true, true, false}
	fmt.Println(r)

	s := []struct {
		x int
		y bool
	}{
		{1, true},
		{2, false},
		{5, true},
		{7, false},
	}
	fmt.Println((s))
}
