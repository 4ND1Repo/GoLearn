package main

import "fmt"

type stc struct {
	x int
	y float64
}

func main() {
	a, b := 200.01, 13

	fmt.Printf("%T \t: %v\n", a, a)
	fmt.Printf("%T \t\t: %v\n", b, b)

	// struct implement
	// struct like class but can't change structure
	p := stc{x: b, y: a}
	fmt.Printf("%T \t: %v\n", p, p)
}
