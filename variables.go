package main

import "fmt"

func main() {

	// explicit declaration and initialization of variable string
	var a = "initial"
	fmt.Println(a)
	// explicit declaration and initialization of variables int
	var b, c int = 1, 2
	fmt.Println(b, c)
	// explicit declaration and initialiization of variable bool
	var d = true
	fmt.Println(d)
	// explicit declaration of variable e tyoe int, since
	// not initialized default value is zero
	var e int
	fmt.Println(e)
	// := sytntax is shorthand for declaring and initialising a variable
	// var f string = "short"
	f := "short"
	fmt.Println(f)
}
