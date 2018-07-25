package main

import "fmt"

func main() {
	var i = [5]int{1, 2, 3, 4}
	var f = [4]float64{1.2, 1.3, 1.4}
	var s = [2]string{"Open", "Shift"}

	fmt.Printf("%T \t\t: %v\n", i, i)
	fmt.Printf("%T \t: %v\n", f, f)
	fmt.Printf("%T \t: %v\n", s, s)
	fmt.Printf("%T \t\t: %v\n", s[1], s[1])

	// set array with make
	var m = make([]int, 5)
	fmt.Printf("%T \t\t: %v\n", m, m)
}
